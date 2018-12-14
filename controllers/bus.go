package controllers

import (
	"net/http"

	"github.com/tockn/tamabus/domain"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/tockn/tamabus/models"
)

type BusController struct {
	DB *sqlx.DB
}

func (controller *BusController) GetBuses(c *gin.Context) {
	buses, err := models.GetAll(controller.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, buses)
}

func (controller *BusController) PostGPS(c *gin.Context) {

	busID, err := paramID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	var bus domain.Bus
	if err := c.BindJSON(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	resBus, err := models.InsertLogByID(controller.DB, busID, &bus)
	c.JSON(http.StatusOK, resBus)
}
