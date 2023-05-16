package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerPort int
}

var config *Config

func GetConfig() *Config {
	if config != nil {
		return config
	}

	serverPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))

	if err != nil {
		log.Panic("serverPort was not be processed")
	}

	config = &Config{
		ServerPort: serverPort,
	}
	return config
}
