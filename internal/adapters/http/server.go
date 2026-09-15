package http

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router          *gin.Engine
	greetingHandler *GreetingHandler
}

func NewServer(greetingHandler *GreetingHandler) *Server {
	s := &Server{
		router:          gin.Default(),
		greetingHandler: greetingHandler,
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

	v1.GET("/greeting", s.greetingHandler.Get)
}
