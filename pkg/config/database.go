package config

import (
	"os"
	"strconv"
	"time"
)

type DatabaseMongodbConfig struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	TimeOut  time.Duration
}

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func DatabaseMongodb() DatabaseMongodbConfig {
	return DatabaseMongodbConfig{
		Host:     os.Getenv("MONGODB_HOST"),
		Port:     os.Getenv("MONGODB_PORT"),
		Database: os.Getenv("MONGODB_NAME"),
		User:     os.Getenv("MONGODB_USERNAME"),
		Password: os.Getenv("MONGODB_PASSWORD"),
		TimeOut:  time.Duration(ConvertInt("APP_TIMEOUT")) * time.Second,
	}
}
