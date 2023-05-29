package server

import (
	"fmt"

	"github.com/bperezgo/go-template/shared/platform/logger"
	"github.com/bperezgo/go-template/shared/platform/middlewares"
	"github.com/gin-gonic/gin"
)

type HandlerMethod int64

const (
	GET HandlerMethod = iota
	Any
	DELETE
	HEAD
	POST
	OPTIONS
	PATCH
	PUT
)

func (s HandlerMethod) String() string {
	switch s {
	case GET:
		return "GET"
	case Any:
		return "Any"
	case DELETE:
		return "DELETE"
	case HEAD:
		return "HEAD"
	case POST:
		return "POST"
	case OPTIONS:
		return "OPTIONS"
	case PATCH:
		return "PATCH"
	case PUT:
		return "PUT"
	}
	return "unknown"
}

type Handler interface {
	GetMethod() HandlerMethod
	GetPath() string
	Function(c *gin.Context)
}

type Server struct {
	engine *gin.Engine

	handlers []Handler
}

func NewServer(handlers ...Handler) *Server {
	r := gin.Default()

	l := logger.GetLogger()

	loggingMiddleware := middlewares.NewLoggingMiddleware(l)

	r.Use(loggingMiddleware.Handle)
	r.Use(gin.Recovery())

	defineGinHandlers(r, handlers)

	return &Server{
		engine:   r,
		handlers: handlers,
	}
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	return s.engine.Run(addr)
}

var mapMethods = make(map[HandlerMethod]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes)

func defineGinHandlers(engine *gin.Engine, handlers []Handler) {
	mapMethods[GET] = engine.GET
	mapMethods[Any] = engine.Any
	mapMethods[DELETE] = engine.DELETE
	mapMethods[HEAD] = engine.HEAD
	mapMethods[POST] = engine.POST
	mapMethods[OPTIONS] = engine.OPTIONS
	mapMethods[PATCH] = engine.PATCH
	mapMethods[PUT] = engine.PUT
	for _, handler := range handlers {
		mapMethods[handler.GetMethod()](handler.GetPath(), handler.Function)
	}
}
