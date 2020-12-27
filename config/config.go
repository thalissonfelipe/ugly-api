package config

import (
	"log"
	"os"
	"strconv"
)

// Config stores API and MongoDB settings
type Config struct {
	API *APIConfig
	DB  *DBConfig
}

// APIConfig stores API settings
type APIConfig struct {
	Host          string
	Port          string
	JWTKey        []byte
	JWTExpireTime int
}

// DBConfig stores MongoDB settings
type DBConfig struct {
	Host           string
	Port           string
	Username       string
	Password       string
	DatabaseName   string
	CollectionName string
}

// MyConfig is a global variable that stores all necessary configurations
var MyConfig *Config

// GetConfig return all config from .env file
func GetConfig() *Config {
	jwtExpireTime, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_TIME"))
	if err != nil {
		log.Fatal(err)
	}
	return &Config{
		API: &APIConfig{
			Host:          os.Getenv("API_HOST"),
			Port:          os.Getenv("API_PORT"),
			JWTKey:        []byte(os.Getenv("JWT_KEY")),
			JWTExpireTime: jwtExpireTime,
		},
		DB: &DBConfig{
			Host:           os.Getenv("DB_HOST"),
			Port:           os.Getenv("DB_PORT"),
			Username:       os.Getenv("DB_USERNAME"),
			Password:       os.Getenv("DB_PASSWORD"),
			DatabaseName:   os.Getenv("DB_DATABASE_NAME"),
			CollectionName: os.Getenv("DB_COLLECTION_NAME"),
		},
	}
}
