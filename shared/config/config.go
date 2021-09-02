package config

import (
	"context"
	"fmt"
	"net/http"

	connDsn "article-web-service/shared/database/postgres"
	log "article-web-service/shared/log/app"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	DevEnvMode  = "development"
	ProdEnvMode = "production"
	TestEnvMode = "testing"
)

func InitConfig() {
	viper.AddConfigPath("shared/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func InitServer() (*gin.Engine, *http.Server) {
	host := viper.GetString("application.server.host")
	port := viper.GetInt("application.server.port")
	env := viper.GetString("application.env")

	switch env {
	case TestEnvMode:
		gin.SetMode(gin.TestMode)
	case ProdEnvMode:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	return r, srv
}

func Start(srv *http.Server) {
	log.Info(nil, "Server is being start")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(&logrus.Fields{"error": err.Error()}, "Error while start server")
	}
}

func Shutdown(srv *http.Server, ctx context.Context) {
	if err := srv.Shutdown(ctx); err != nil {
		log.Error(&logrus.Fields{"error": err.Error()}, "Error while shutdown server")
	}
}

func InitAppLogger(hooks ...logrus.Hook) {
	log.Init(func(l *logrus.Logger) {
		for _, hook := range hooks {
			l.AddHook(hook)
		}
	})
}

func InitPostgresDatabase(database string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(connDsn.ConnectionDsn(database)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while init postgres database")
	}

	return db
}

func InitLocalizerI18n() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("shared/config/localize/en.toml")
	bundle.MustLoadMessageFile("shared/config/localize/id.toml")

	return bundle
}

func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("application.resources.redis.host"), viper.GetInt("application.resources.redis.port")),
		Password: viper.GetString("application.resources.redis.password"),
		DB:       viper.GetInt("application.resources.redis.db"),
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while init redis")
	}

	return client
}

func InitElasticsearch() *elastic.Client {
	host := viper.GetString("application.resources.elasticsearch.host")
	port := viper.GetInt("application.resources.elasticsearch.port")

	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("%s:%d", host, port)),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	)

	if err != nil {
		log.Panic(&logrus.Fields{"error": err.Error()}, "Error while init elasticsearch")
	}

	return client
}
