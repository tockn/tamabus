package handler

import "github.com/gin-gonic/gin"

type AppHandler interface {
	BusHandler
}

type BusHandler interface {
	GetData(*gin.Context)
}

type busHandler struct {
	u usecase.BusUseCase
}

func NewBusHandler(u usecase.BusHandler) BusHandler {
	return &busHandler{u}
}

func (b *busHandler) GetData(*gin.Context) {
}
