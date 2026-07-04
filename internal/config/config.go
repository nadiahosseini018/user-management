package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}

	cfg := &Config{
		AppPort:    os.Getenv("APP_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}

	fmt.Printf("%+v\n", cfg)

	return cfg
	// return &Config{
	// 	AppPort: os.Getenv("APP_PORT"),

	// 	DBHost:     os.Getenv("DB_HOST"),
	// 	DBPort:     os.Getenv("DB_PORT"),
	// 	DBUser:     os.Getenv("DB_USER"),
	// 	DBPassword: os.Getenv("DB_PASSWORD"),
	// 	DBName:     os.Getenv("DB_NAME"),
	// }
}
