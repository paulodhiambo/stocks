package router

import (
	"github.com/gin-gonic/gin"
	"stocks/config"
)

func Init(init *config.Initialization) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			stocks := v1.Group("/stock")
			stocks.GET("/", init.StockCtrl.GetAllStocks)
			stocks.GET("/:id", init.StockCtrl.GetStockById)
			stocks.POST("/", init.StockCtrl.AddStock)
			stocks.DELETE("/:id", init.StockCtrl.DeleteStock)
		}

	}
	return router
}
