package application

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddress string
	ServerPort   uint16
	Password     string
}

func LoadConfig() Config {
	// load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}

	redisAddress := os.Getenv("REDIS_ADDR")
	password := os.Getenv("PASSWORD")

	// set the config object
	cfg := Config{
		RedisAddress: redisAddress,
		ServerPort:   3000,
		Password:     password,
	}

	if redisAddr, exists := os.LookupEnv("REDIS_ADDR"); exists {
		cfg.RedisAddress = redisAddr
	}

	if serverPort, exists := os.LookupEnv("SERVER_PORT"); exists {
		if port, err := strconv.ParseUint(serverPort, 10, 16); err == nil {
			cfg.ServerPort = uint16(port)
		}
	}

	return cfg
}
