package main

import (
	"article-web-service/read-service/router"
	"article-web-service/shared/config"
	"article-web-service/shared/domains"
	"article-web-service/shared/middleware"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	hooks "article-web-service/shared/log"
	log "article-web-service/shared/log/app"

	"github.com/spf13/viper"
)

var (
	serviceName string
	env         string
)

func init() {
	config.InitConfig()

	serviceName = "read-service"
	env = viper.GetString("application.env")

	config.InitAppLogger(hooks.NewEnvFieldHook(serviceName, env))
}

func closeAllInstance(fn func()) {
	log.Info(nil, "Closing all instance")

	fn()

	log.Info(nil, "All instance are closed")
}

func main() {
	var appContext domains.Context

	bundle := config.InitLocalizerI18n()
	esClient := config.InitElasticsearch()

	r, srv := config.InitServer()
	r.Use(middleware.CORS())
	r.Use(middleware.Localizer(&appContext, bundle))

	appContext.R = r
	appContext.ESClient = esClient

	router.Configure(&appContext)

	closeInstance := func() {}

	go config.Start(srv)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info(nil, "Shutdown server")

	closeAllInstance(closeInstance)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	config.Shutdown(srv, ctx)
	<-ctx.Done()
	log.Info(nil, "Server has shutdown")
}
