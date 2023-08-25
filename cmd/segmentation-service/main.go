package main

import (
	"context"
	"os"
	"os/signal"
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/handler"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	//TODO: init config - viper
	//TODO: init logger - logrus
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("config init error: %s", err.Error())
	}

	//TODO: init db - postgres in docker

	//TODO: init router - gin

	handlers := new(handler.Handler)

	srv := &segmentation_service.Server{}
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("an error occurred while running the server: %s", err.Error())
		}
	}()

	logrus.Print("Application started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Application finished")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("server shutting down error: %s", err.Error())
	}
	//TODO: run server
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
