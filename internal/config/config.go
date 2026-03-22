package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT string
	APP_ENV string
	APP_TIMEZONE string
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	JWT_SECRET string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load .env files")
	}

	log.Println("Config loaded")

	return &Config{
		APP_PORT: os.Getenv("APP_PORT"),
		APP_ENV: os.Getenv("APP_ENV"),
		APP_TIMEZONE: os.Getenv("APP_TIMEZONE"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME: os.Getenv("DB_NAME"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}, nil
}