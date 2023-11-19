package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bperezgo/go-template/shared/platform/handlertypes"
	"github.com/bperezgo/go-template/shared/platform/logger"
	"github.com/gin-gonic/gin"
)

type JsonHandler struct {
	logger *logger.Logger
}

func NewJsonHandler() *JsonHandler {
	return &JsonHandler{
		logger: logger.GetLogger(),
	}
}

func (h *JsonHandler) Adapt(handler Handler) func(c *gin.Context) {
	return func(c *gin.Context) {
		method := handler.GetMethod()
		request := handlertypes.Request{}
		var body interface{}
		if method == POST {
			if err := c.ShouldBindJSON(&body); err != nil {
				h.logger.Error(logger.LogInput{
					Action: "JsonHandler.Adapt",
					State:  logger.FAILED,
					Error: &logger.Error{
						Message: err.Error(),
					},
				})
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})
				return
			}

			headers := handlertypes.Headers{}

			if err := c.ShouldBindHeader(&headers); err != nil {
				h.logger.Error(logger.LogInput{
					Action: "JsonHandler.Adapt",
					State:  logger.FAILED,
					Error: &logger.Error{
						Message: err.Error(),
					},
				})
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})
				return
			}

			b, err := json.Marshal(body)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Invalid request body",
				})

				return
			}

			request.Body = b

			request.Headers = headers
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
