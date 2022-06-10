package postgres

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode)),
		&gorm.Config{})
	if err != nil {
		logrus.Println(err.Error())
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		logrus.Fatal("Failed to setup sql database:", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		logrus.Println(err.Error())
		return nil, err
	}

	return gormDB, nil
}
