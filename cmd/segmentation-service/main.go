package main

import (
	"context"
	"os"
	"os/signal"
	segmentation_service "segmentation-service"
	"segmentation-service/pkg/handler"
	"segmentation-service/pkg/repository"
	"segmentation-service/pkg/service"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"segmentation-service/docs"
)

func main() {
	docs.SwaggerInfo.Title = "Dynamic Segmentation Service"
	docs.SwaggerInfo.Description = "API Server for Dynamic User Segmentation App"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "127.0.0.1:8081"
	docs.SwaggerInfo.BasePath = "/"

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("config init error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("env reading error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("PG_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

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

	if err := db.Close(); err != nil {
		logrus.Errorf("db connection close error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
