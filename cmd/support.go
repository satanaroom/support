package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"support"
	"support/pkg/handler"
	"support/pkg/repository"
	"support/pkg/service"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("[cfg] initialization config error: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("[env] loading env error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("[db] initialization database error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(support.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("[srv] running sever error: %s", err.Error())
		}
	}()
	logrus.Print("SupportApp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("[srv] server shutting down error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("[db] db close error: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
