package controller

import (
	"github.com/gin-gonic/gin"
	"stocks/app/service"
)

type StockController interface {
	GetAllStocks(c *gin.Context)
	AddStock(c *gin.Context)
	GetStockById(c *gin.Context)
	DeleteStock(c *gin.Context)
}

type StockControllerImpl struct {
	svc service.StockService
}

func (s StockControllerImpl) GetAllStocks(c *gin.Context) {
	s.svc.GetAllStocks(c)
}

func (s StockControllerImpl) AddStock(c *gin.Context) {
	s.svc.AddStock(c)
}

func (s StockControllerImpl) GetStockById(c *gin.Context) {
	s.svc.GetStockById(c)
}

func (s StockControllerImpl) DeleteStock(c *gin.Context) {
	s.svc.DeleteStock(c)
}

func StockControllerInit(stockService service.StockService) *StockControllerImpl {
	return &StockControllerImpl{
		svc: stockService,
	}
}
