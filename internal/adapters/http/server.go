package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router      *gin.Engine
	userHandler *UserHandler
}

func NewServer(userHandler *UserHandler) *Server {
	s := &Server{
		router:      gin.Default(),
		userHandler: userHandler,
	}
	s.registerRoutes()
	return s
}

func (s *Server) Run(port string) error {
	return s.router.Run(fmt.Sprintf(":%s", port))
}

func (s *Server) registerRoutes() {
	api := s.router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/users/:id", s.userHandler.Get)
	v1.POST("/users", s.userHandler.Create)
	v1.PUT("/users/:id", s.userHandler.Update)
	v1.DELETE("/users/:id", s.userHandler.Delete)
}
