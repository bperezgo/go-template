package middlewares

import (
	"github.com/bperezgo/testing-ai/shared/platform/logger"
	"github.com/gin-gonic/gin"
)

type LoggingMiddleware struct {
	logger *logger.Logger
}

func NewLoggingMiddleware(logger *logger.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

func (m *LoggingMiddleware) Handle(c *gin.Context) {
	m.logger.Info(logger.LogInput{
		Action: "REQUEST",
		State:  "SUCCESS",
	})
	c.Next()
	m.logger.Info(logger.LogInput{
		Action: "RESPONSE",
		State:  "SUCCESS",
	})
}
