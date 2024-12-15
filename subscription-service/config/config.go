package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	DBUrl string
	RedisUrl  string
	JwtSecret string
}

func LoadConfig() Config {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	return Config{
		Port:  os.Getenv("PORT"),
		DBUrl: os.Getenv("DATABASE_URL"),
		RedisUrl:  os.Getenv("REDIS_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
}
