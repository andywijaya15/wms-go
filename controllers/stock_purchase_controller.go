package controllers

import (
	"net/http"
	"time"
	"wms-go/models"

	"github.com/gin-gonic/gin"
)

func DeleteStockPurchase(c *gin.Context) {
	var stockPurchase []models.StockPurchase
	cOrderId := c.PostForm("c_order_id")

	err := models.DB.Where(&stockPurchase, "c_order_id = ?", cOrderId).
		Update("deleted_at", time.Now()).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Stock Purchase Deleted Successfully"})
}
