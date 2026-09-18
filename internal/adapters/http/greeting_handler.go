package http

import (
	nethttp "net/http"

	"github.com/gin-gonic/gin"
)

type greetingService interface {
	Get() string
}

type GreetingHandler struct {
	service greetingService
}

func NewGreetingHandler(service greetingService) *GreetingHandler {
	return &GreetingHandler{service: service}
}

func (h *GreetingHandler) Get(ctx *gin.Context) {
	message := h.service.Get()
	ctx.JSON(nethttp.StatusOK, newGreetingResponse(message))
}
