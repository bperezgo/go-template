package handlers

import (
	"github.com/bperezgo/go-template/shared/platform/server"
)

func GetHandlers() []server.Handler {
	getHealth := NewGetHealthHandler()
	return []server.Handler{
		getHealth,
	}
}
