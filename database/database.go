package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     string
	DBname   string
	User     string
	Password string
	SSLMode  string
}

func BuildDBConfig() *DBConfig {
	if err := godotenv.Load("./.env"); err != nil {
		panic(err.Error())
	}

	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBname:   os.Getenv("DB_DBNAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
	}
	return &dbConfig
}

func DbConnect(dbConfig *DBConfig) string {
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Bangkok",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBname,
		dbConfig.Port,
		dbConfig.SSLMode,
	)
}
