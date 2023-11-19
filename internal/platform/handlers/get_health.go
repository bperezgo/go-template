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

type GetHealthHandler struct {
	handler.BasisHandler
}

func NewGetHealthHandler() *GetHealthHandler {
	return &GetHealthHandler{
		BasisHandler: handler.BasisHandler{
			HandlerMethod: handler.GET,
			Path:          "/health",
		},
	}
}

func (h *GetHealthHandler) Function(req handlertypes.Request) handlertypes.Response {
	return handlertypes.Response{
		Body: GetHealthResponse{
			Message: "Server is up and running",
		},
		HttpStatus: http.StatusOK,
	}
}
