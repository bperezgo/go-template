package handlers

import (
	"github.com/bperezgo/go-template/shared/platform/handler"
)

func GetHandlers() []handler.Handler {
	getHealth := NewGetHealthHandler()
	createVideos := NewCreateVideoHandler()
	return []handler.Handler{
		getHealth,
		createVideos,
	}
}
