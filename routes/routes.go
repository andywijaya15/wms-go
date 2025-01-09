package routes

import (
	"wms-go/controllers"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.BestSpeed))

	v1 := router.Group("/v1")
	{
		v1.DELETE("/delete-stock-purchase", controllers.DeleteStockPurchase)
		v1.GET("/get-factory/:id", controllers.GetFactory)
	}

	return router
}
