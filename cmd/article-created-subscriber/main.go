package main

import (
	"article-web-service/article-created-subscriber/models"
	"article-web-service/shared/config"
	"article-web-service/shared/utils"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	hooks "article-web-service/shared/log"
	log "article-web-service/shared/log/app"

	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	serviceName string
	env         string
)

func init() {
	config.InitConfig()

	serviceName = "article-created-subscriber"
	env = viper.GetString("application.env")

	config.InitAppLogger(hooks.NewEnvFieldHook(serviceName, env))
}

func closeAllInstance(fn func()) {
	log.Info(nil, "Closing all instance")

	fn()

	log.Info(nil, "All instance are closed")
}

func main() {
	redisClient := config.InitRedis()
	ctx := context.Background()
	topic := redisClient.Subscribe(ctx, "article_created")
	channel := topic.Channel()
	esClient := config.InitElasticsearch()

	for msg := range channel {
		fmt.Println(msg.Payload)

		article := models.Article{}
		if err := utils.Unmarshal([]byte(msg.Payload), &article); err != nil {
			log.Error(&logrus.Fields{"error": err.Error()}, "Error when decode payload")
		}

		if err := indexToES(esClient, &article); err != nil {
			log.Error(&logrus.Fields{"error": err.Error()}, "Error when index to ES")
		}
	}

	closeInstance := func() {}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(nil, "Shutdown server")

	closeAllInstance(closeInstance)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	<-ctx.Done()
	log.Info(nil, "Server has shutdown")
}

func indexToES(client *elastic.Client, article *models.Article) error {
	data, _ := utils.Marshal(article)
	id := fmt.Sprintf("%d", article.ID)
	_, err := client.Index().Index("articles").Id(id).BodyJson(string(data)).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
