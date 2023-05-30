package handlers

import (
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handler"
)

type CreateVideoResponse struct {
	Message string `json:"message"`
}

type CreateVideoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateVideoHandler struct {
}

func NewCreateVideoHandler() *CreateVideoHandler {
	return &CreateVideoHandler{}
}

func (h *CreateVideoHandler) GetMethod() handler.HandlerMethod {
	return handler.POST
}

func (h *CreateVideoHandler) GetPath() string {
	return "/videos"
}

func (h *CreateVideoHandler) Function(req handler.Request) handler.Response {
	return handler.Response{
		Body: CreateVideoResponse{
			Message: "Video created",
		},
		HttpStatus: http.StatusCreated,
	}
}

func (h *CreateVideoHandler) GetEmptyRequest() handler.Request {
	return handler.Request{
		Body: CreateVideoRequest{},
	}
}
