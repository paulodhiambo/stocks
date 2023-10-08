package config

import (
	"stocks/app/controller"
	"stocks/app/repository"
	"stocks/app/service"
)

type Initialization struct {
	stockRepo repository.StockRepository
	stockSvc  service.StockService
	StockCtrl controller.StockController
}

func NewInitialization(stockRepo repository.StockRepository,
	stockService service.StockService,
	stockCtrl controller.StockController,
) *Initialization {
	return &Initialization{
		stockRepo: stockRepo,
		stockSvc:  stockService,
		StockCtrl: stockCtrl,
	}
}
