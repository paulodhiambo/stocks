package service

import (
	"github.com/gin-gonic/gin"
	"stocks/app/repository"
)

type StockService interface {
	GetAllStocks(c *gin.Context)
	GetStockById(c *gin.Context)
	AddStock(c *gin.Context)
	DeleteStock(c *gin.Context)
}

type StockServiceImpl struct {
	stockRepository repository.StockRepository
}

func (s StockServiceImpl) GetAllStocks(c *gin.Context) {

}

func (s StockServiceImpl) GetStockById(c *gin.Context) {

}

func (s StockServiceImpl) AddStock(c *gin.Context) {

}

func (s StockServiceImpl) DeleteStock(c *gin.Context) {

}

func StockServiceInit(stockRepository repository.StockRepository) *StockServiceImpl {
	return &StockServiceImpl{
		stockRepository: stockRepository,
	}
}
