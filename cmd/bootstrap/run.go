package bootstrap

import (
	"github.com/bperezgo/testing-ai/internal/platform/handlers"
	"github.com/bperezgo/testing-ai/shared/platform/logger"
	"github.com/bperezgo/testing-ai/shared/platform/server"
)

func Run() error {
	logger.InitLogger()
	allHandlers := handlers.GetHandlers()
	srv := server.NewServer(allHandlers...)
	return srv.Start(8000)
}
