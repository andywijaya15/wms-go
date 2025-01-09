package controllers

import (
	"net/http"
	"wms-go/models"

	"github.com/gin-gonic/gin"
)

func GetFactory(c *gin.Context) {
	var factory []models.Factory
	id := c.Query("id")

	err := models.DB.First(&factory, "id = ?", id).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock Purchase Deleted Successfully", "data": factory})
}
