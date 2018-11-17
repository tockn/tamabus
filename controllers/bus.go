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

func (controller *BusController) PostGPS(c *gin.Context) {

	busID, err := paramID(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, "internal server error")
		return
	}

	var bus models.Bus
	if err := c.BindJSON(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	resBus, err := models.UpdatePosByID(busID, bus)
	c.JSON(http.StatusOK, resBus)
}
