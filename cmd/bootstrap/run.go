package bootstrap

import (
	"log"

	"github.com/bperezgo/testing-ai/internal/platform/config"
	"github.com/bperezgo/testing-ai/internal/platform/handlers"
	"github.com/bperezgo/testing-ai/shared/platform/logger"
	"github.com/bperezgo/testing-ai/shared/platform/server"
	"github.com/joho/godotenv"
)

func Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := config.GetConfig()
	logger.InitLogger()
	allHandlers := handlers.GetHandlers()
	srv := server.NewServer(allHandlers...)
	return srv.Start(c.ServerPort)
}
