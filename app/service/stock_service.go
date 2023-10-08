package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"stocks/app/constant"
	"stocks/app/domain/dao"
	"stocks/app/pkg"
	"stocks/app/repository"
	"strconv"
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
	defer pkg.PanicHandler(c)
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Errorf("Invalid page")
		pkg.PanicException(constant.InvalidRequest)
	}
	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if err != nil {
		log.Errorf("Invalid page size")
		pkg.PanicException(constant.InvalidRequest)
	}
	data, err := s.stockRepository.FindAllStocks(page, pageSize)
	if err != nil {
		log.Errorf("Unable to retrieve stocks")
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))
}

func (s StockServiceImpl) GetStockById(c *gin.Context) {
	defer pkg.PanicHandler(c)
	stockId, _ := strconv.Atoi(c.Param("id"))
	data, err := s.stockRepository.FindStockById(stockId)
	if err != nil {
		log.Errorf("Stock with the given Id not found %s", err)
		pkg.PanicException(constant.DataNotFound)
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, data))

}

func (s StockServiceImpl) AddStock(c *gin.Context) {
	defer pkg.PanicHandler(c)

	var request dao.Stock

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Errorf("Invalid data %s", err)
		pkg.PanicException(constant.InvalidRequest)
	}
	data, err := s.stockRepository.SaveStock(&request)
	if err != nil {
		log.Errorf("Unable to save stock information, please try again %s", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, data))

}

func (s StockServiceImpl) DeleteStock(c *gin.Context) {
	defer pkg.PanicHandler(c)
	stockId, _ := strconv.Atoi(c.Param("id"))
	err := s.stockRepository.DeleteStockById(stockId)
	if err != nil {
		log.Errorf("Unable to delete stock with the given Id %s", err)
		pkg.PanicException(constant.UnknownError)
	}
	c.JSON(http.StatusNoContent, pkg.BuildResponse(constant.Success, pkg.Null()))
}

func StockServiceInit(stockRepository repository.StockRepository) *StockServiceImpl {
	return &StockServiceImpl{
		stockRepository: stockRepository,
	}
}
