package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return &Server{
		engine: r,
	}
}

func (s *Server) Start(port int) error {
	addr := fmt.Sprintf("localhost:%d", port)
	return s.engine.Run(addr)
}
