package config

import (
	"fmt"
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
	JWTExpiration int64
	JWTSecret     string
}

var Envs = initConfig()

func initConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return Config{}
	}

	config := Config{
		PublicHost:    getEnv("PUBLIC_HOST"),
		Port:          getEnv("PORT"),
		DBUser:        getEnv("DB_USER"),
		DBPassword:    getEnv("DB_PASSWORD"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnv("DB_HOST"), getEnv("DB_PORT")),
		DBName:        getEnv("DB_NAME"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION"),
		JWTSecret:     getEnv("JWT_SECRET"),
	}

	return config
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Error: %s environment variable not set", key)
	}
	return value
}

func getEnvAsInt(key string) int64 {
	valueStr := getEnv(key)
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		log.Fatalf("Error: %s environment variable must be an integer", key)
	}
	return value
}
