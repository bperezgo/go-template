package handlers

import (
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/handlertypes"
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

func (h *CreateVideoHandler) Function(req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: CreateVideoResponse{
			Message: "Video created",
		},
		HttpStatus: http.StatusCreated,
	}
}

func (h *CreateVideoHandler) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{
		Body: CreateVideoRequest{},
	}
}
