package main

import (
	"context"
	"github.com/AssylzhanZharzhanov/task-management-app/internal/app"
	"github.com/AssylzhanZharzhanov/task-management-app/pkg/db/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error in reading env file: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
		SSLMode:  viper.GetString("database.ssl_mode"),
	})
	if err != nil {
		log.Fatalf("Error in starting database: %s", err.Error())
	}

	port := viper.GetString("server.port")

	srv := app.NewApp(db, port)
	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalf("Error in starting server: %s", err.Error())
		}
	}()
	logrus.Print("Server started at " + port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("can not setup sql database")
	}

	if err := sqlDB.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
