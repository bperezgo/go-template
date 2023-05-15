package handlers

import (
	"github.com/bperezgo/testing-ai/shared/platform/server"
)

func GetHandlers() []server.Handler {
	getHealth := NewGetHealthHandler()
	return []server.Handler{
		getHealth,
	}
}
