package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetHealthHandler struct{}

func NewGetHealthHandler() *GetHealthHandler {
	return &GetHealthHandler{}
}

func (h *GetHealthHandler) GetMethod() string {
	return "GET"
}

func (h *GetHealthHandler) GetPath() string {
	return "/health"
}

func (h *GetHealthHandler) Function(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Server is up and running",
	})
}
