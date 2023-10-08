//go:build wireinject
// +build wireinject

// go:build wireinject
package config

import (
	"github.com/google/wire"
	"stocks/app/controller"
	"stocks/app/repository"
	"stocks/app/service"
)

var db = wire.NewSet(ConnectToDB)

var stockServiceSet = wire.NewSet(service.StockServiceInit,
	wire.Bind(new(service.StockService), new(*service.StockServiceImpl)),
)

var stockRepoSet = wire.NewSet(repository.StockRepositoryInit,
	wire.Bind(new(repository.StockRepository), new(*repository.StockRepositoryImpl)),
)

var stockCtrlSet = wire.NewSet(controller.StockControllerInit,
	wire.Bind(new(controller.StockController), new(*controller.StockControllerImpl)),
)

func Init() *Initialization {
	wire.Build(NewInitialization, db, stockCtrlSet, stockServiceSet, stockRepoSet)
	return nil
}
