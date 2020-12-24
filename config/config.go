package config

import (
	"os"
)

// Config stores API and MongoDB settings
type Config struct {
	API *APIConfig
	DB  *DBConfig
}

// APIConfig stores API settings
type APIConfig struct {
	Host string
	Port string
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
	return &Config{
		API: &APIConfig{
			Host: os.Getenv("API_HOST"),
			Port: os.Getenv("API_PORT"),
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
