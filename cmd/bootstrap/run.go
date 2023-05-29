package bootstrap

import (
	"log"

	"github.com/bperezgo/go-template/internal/platform/config"
	"github.com/bperezgo/go-template/internal/platform/handlers"
	"github.com/bperezgo/go-template/shared/platform/logger"
	"github.com/bperezgo/go-template/shared/platform/server"
)

func Run() error {
	err := config.InitConfig()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}

	c := config.GetConfig()
	logger.InitLogger()
	allHandlers := handlers.GetHandlers()
	srv := server.NewServer(allHandlers...)
	return srv.Start(c.ServerPort)
}
