package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func paramID(c *gin.Context, key string) (int64, error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}
