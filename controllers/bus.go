package controllers

import (
	"encoding/base64"
	"github.com/tockn/tamabus/domain"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/tockn/tamabus/models"
)

type BusController struct {
	DB     *sqlx.DB
	Logger *log.Logger
}

func (controller *BusController) GetBuses(c *gin.Context) {
	buses, err := models.GetAll(controller.DB)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, buses)
}

func (controller *BusController) PostGPS(c *gin.Context) {
	var dbus domain.Bus
	if err := c.BindJSON(&dbus); err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	resBus, err := models.InsertLog(controller.DB, &dbus)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
	c.JSON(http.StatusOK, resBus)
}

func (controller *BusController) PostImage(c *gin.Context) {
	var img domain.BusImage
	if err := c.BindJSON(&img); err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	bin, err := base64.StdEncoding.DecodeString(img.Base64)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusBadRequest, "could not decode from base64")
		return
	}

	mi := models.BusImage{
		BusID: img.BusID,
		Body: string(bin),
	}

	err = mi.Insert(controller.DB)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
}
