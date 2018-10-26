package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tockn/tamabus/presenter/api/handler"
)

func NewRouter(h *handler.AppHandler) {
	router := gin.New()

	router.GET("/api/bus", h.GetBus)

	return router
}
