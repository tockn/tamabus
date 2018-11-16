package controllers

import (
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/tockn/tamabus/models"
)

type BusController struct {
	DB *sqlx.DB
}

func (controller *BusController) GetBuses(c *gin.Context) {
	buses, err := models.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, buses)
}
