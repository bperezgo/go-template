package handlers

import (
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/logger"
)

type CreateVideoResponse struct {
	Message string `json:"message"`
}

type CreateVideoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateVideoHandler struct {
	logger logger.Logger
}

func NewCreateVideoHandler() *CreateVideoHandler {
	return &CreateVideoHandler{
		logger: *logger.GetLogger(),
	}
}

func (h *CreateVideoHandler) GetMethod() handler.HandlerMethod {
	return handler.POST
}

func (h *CreateVideoHandler) GetPath() string {
	return "/videos"
}

func (h *CreateVideoHandler) Function(req handler.Request) handler.Response {
	h.logger.Info(logger.LogInput{
		Action:  "createVideo",
		State:   logger.SUCCESS,
		Message: "Video created",
	})
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
