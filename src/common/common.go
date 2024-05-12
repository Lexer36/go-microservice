package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	// LogFilename   string
	// LogMaxSize    int
	// LogMaxBackups int
	// LogMaxAge     int

	DbAddr     string
	DbName     string
	DbUser     string
	DbPassword string
}

var (
	Config *Configuration
)

func LoadConfig() error {
	// Загрузка переменных среды
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("ошибка загрузки файла .env")
	}

	Config = &Configuration{
		DbAddr:     os.Getenv("DB_ADDR"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}

	// устанавливаем логгер
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/log.log",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     1,
	})
	log.SetLevel(log.DebugLevel)

	// log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

	return nil
}
