package handlers

import (
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/handlertypes"
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

func (h *GetHealthHandler) Function(req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: GetHealthResponse{
			Message: "Server is up and running",
		},
		HttpStatus: http.StatusOK,
	}
}

func (h *GetHealthHandler) GetEmptyRequest() handlertypes.Request {
	return handlertypes.Request{}
}
