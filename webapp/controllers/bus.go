package controllers

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/tockn/tamabus/webapp/domain"

	"github.com/jmoiron/sqlx"

	"github.com/gin-gonic/gin"
	"github.com/tockn/tamabus/webapp/models"
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

	fileName, err := decode(img.Base64, img.FileType)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusBadRequest, "could not decode from base64")
		return
	}

	mi := models.BusImage{
		BusID: img.BusID,
		Path:  string(fileName),
	}

	err = mi.Insert(controller.DB)
	if err != nil {
		controller.Logger.Println(err)
		c.JSON(http.StatusInternalServerError, "internal server error")
		return
	}
}

func decode(body, fileType string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%d.%s", time.Now().Unix(), fileType)
	fullPath := fmt.Sprintf("./congestionCalculator/images/%s", fileName)
	file, err := os.Create(fullPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = file.Write(data)
	return fileName, err
}
