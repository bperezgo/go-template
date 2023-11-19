package handlers

import (
	"github.com/bperezgo/go-template/internal/platform/repositories"
	"github.com/bperezgo/go-template/shared/platform/handler"
)

func GetHandlers() []handler.Handler {
	getHealth := NewGetHealthHandler()
	cvh := NewCreateVideoHandler()

	userRepository := repositories.NewInMemoryUserRepository()
	basicAuth := NewBasicAuthMiddleware(userRepository)
	return []handler.Handler{
		getHealth,
		basicAuth,
		cvh,
	}
}
