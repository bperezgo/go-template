package middlewares

import (
	"github.com/bperezgo/go-template/shared/platform/logger"
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
	// b, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	m.logger.Info(logger.LogInput{
	// 		Action: "REQUEST",
	// 		State:  "FAILED",
	// 		Error: &logger.Error{
	// 			Message: err.Error(),
	// 		},
	// 	})
	// }

	request := &struct {
		Body   interface{}
		Params interface{}
		Query  interface{}
	}{
		// Body:  string(b),
		Query: c.Request.URL.Query(),
	}

	m.logger.Info(logger.LogInput{
		Action: "REQUEST",
		State:  "SUCCESS",
		Http: &logger.LogHttpInput{
			Request: request,
		},
	})
	c.Next()

	m.logger.Info(logger.LogInput{
		Action: "RESPONSE",
		State:  "SUCCESS",
		Http: &logger.LogHttpInput{
			Request: request,
			Response: &struct{ Data interface{} }{
				Data: "",
			},
		},
	})
}
