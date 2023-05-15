package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetMethod() string
	GetPath() string
	Function(c *gin.Context)
}

type Server struct {
	engine *gin.Engine

	handlers []Handler
}

func NewServer(handlers ...Handler) *Server {
	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	for _, handler := range handlers {
		r.GET(handler.GetPath(), handler.Function)
	}

	return &Server{
		engine:   r,
		handlers: handlers,
	}
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	return s.engine.Run(addr)
}

var mapMethods = make(map[string]func(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes)

func defineGinHandlers(engine *gin.Engine, handlers []Handler) *gin.Engine {
	mapMethods["GET"] = engine.GET
	mapMethods["Any"] = engine.Any
	mapMethods["DELETE"] = engine.DELETE
	mapMethods["HEAD"] = engine.HEAD
	mapMethods["POST"] = engine.POST
	mapMethods["OPTIONS"] = engine.OPTIONS
	mapMethods["PATCH"] = engine.PATCH
	mapMethods["PUT"] = engine.PUT
	for i := 1; i < len(handlers); i++ {
		handler := handlers[i]
		mapMethods[handler.GetMethod()](handler.GetPath(), handler.Function)
	}

	return engine
}
