package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonHandler struct{}

func (h *JsonHandler) Adapt(handler Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		method := handler.GetMethod()
		request := handler.GetEmptyRequest()
		if method == POST {
			if err := c.ShouldBindJSON(&request.Body); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})
				return
			}

			response := handler.Function(request)

			c.JSON(http.StatusOK, response.Body)

			return
		}

		if method == GET {
			// TODO: Get query params
			response := handler.Function(request)

			c.JSON(http.StatusOK, response.Body)
			return
		}
	}
}
