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
		v1.GET("/get-warehouse/:id", controllers.GetChangedPurchaseOrders)
		v1.DELETE("/delete-stock-purchase", controllers.DeletePurchaseOrder)
	}

	return router
}
