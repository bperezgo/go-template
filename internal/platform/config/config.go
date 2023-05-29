package config

import (
	"log"
	"os"
	"path"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
}

var config *Config

func InitConfig() error {
	environment, ok := os.LookupEnv("ENVIRONMENT")

	if !ok {
		return godotenv.Load()
	}

	if environment == "local" {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		p := path.Join(cwd, "../.env")
		return godotenv.Load(os.ExpandEnv(p))
	}

	return nil

}

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
