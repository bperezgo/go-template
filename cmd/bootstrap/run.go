package bootstrap

import "github.com/bperezgo/testing-ai/shared/platform/server"

func Run() error {
	srv := server.NewServer()
	return srv.Start(8000)
}
