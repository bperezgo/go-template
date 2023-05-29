package handlers

import (
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handler"
)

type GetHealthResponse struct {
	Message string `json:"message"`
	handler.BaseHandler
}

type GetHealthHandler struct{}

func NewGetHealthHandler() *GetHealthHandler {
	return &GetHealthHandler{}
}

func (h *GetHealthHandler) GetMethod() handler.HandlerMethod {
	return handler.GET
}

func (h *GetHealthHandler) GetPath() string {
	return "/health"
}

func (h *GetHealthHandler) Function(req handler.Request) handler.Response {
	return handler.Response{
		Body: GetHealthResponse{
			Message: "Server is up and running",
		},
		HttpStatus: http.StatusOK,
	}
}

func (h *GetHealthHandler) GetEmptyRequest() handler.Request {
	return handler.Request{}
}
