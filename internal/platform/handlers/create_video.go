package handlers

import (
	"encoding/json"
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
	handler.BasisHandler
}

func NewCreateVideoHandler() *CreateVideoHandler {
	return &CreateVideoHandler{
		BasisHandler: handler.BasisHandler{
			HandlerMethod: handler.POST,
			Path:          "/videos",
		},
	}
}

func (h *CreateVideoHandler) Function(req handlertypes.Request) handlertypes.Response {
	body := CreateVideoRequest{}
	if err := json.Unmarshal(req.Body, &body); err != nil {
		return handlertypes.Response{
			Body:       nil,
			HttpStatus: http.StatusBadRequest,
		}
	}

	return handlertypes.Response{
		Body: CreateVideoResponse{
			Message: "Video created",
		},
		HttpStatus: http.StatusCreated,
	}
}
