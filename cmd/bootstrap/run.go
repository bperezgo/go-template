package bootstrap

import (
	"github.com/bperezgo/testing-ai/internal/platform/handlers"
	"github.com/bperezgo/testing-ai/shared/platform/server"
)

func Run() error {
	allHandlers := handlers.GetHandlers()
	srv := server.NewServer(allHandlers...)
	return srv.Start(8000)
}
