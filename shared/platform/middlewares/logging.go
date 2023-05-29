package middlewares

import (
	"github.com/bperezgo/go-template/shared/platform/handler"
	"github.com/bperezgo/go-template/shared/platform/logger"
)

type LoggerHandler struct {
	handlerFunction handler.Function
	handler         handler.Handler
}

func NewLoggerHandler(handler handler.Handler) *LoggerHandler {
	return &LoggerHandler{
		handlerFunction: decorateWithLogger(handler.Function, logger.GetLogger()),
		handler:         handler,
	}
}

func (h *LoggerHandler) GetMethod() handler.HandlerMethod {
	return h.handler.GetMethod()
}

func (h *LoggerHandler) GetPath() string {
	return h.handler.GetPath()
}

func (h *LoggerHandler) Function(req handler.Request) handler.Response {
	return h.handlerFunction(req)
}

func (h *LoggerHandler) GetEmptyRequest() handler.Request {
	return h.handler.GetEmptyRequest()
}

func decorateWithLogger(handlerFunction handler.Function, l *logger.Logger) handler.Function {
	return func(req handler.Request) (res handler.Response) {
		l.Info(logger.LogInput{
			Action: "REQUEST",
			State:  logger.SUCCESS,
			Http: &logger.LogHttpInput{
				Request: req,
			},
		})

		res = handlerFunction(req)

		l.Info(logger.LogInput{
			Action: "RESPONSE",
			State:  logger.SUCCESS,
			Http: &logger.LogHttpInput{
				Request:  req,
				Response: res,
			},
		})

		return res
	}
}
