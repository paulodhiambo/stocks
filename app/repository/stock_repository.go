package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"stocks/app/domain/dao"
)

type StockRepository interface {
	FindAllStocks() ([]dao.Stock, error)
	FindStockById(id int) (*dao.Stock, error)
	SaveStock(stock *dao.Stock) (*dao.Stock, error)
	DeleteStockById(id int) error
}

type StockRepositoryImpl struct {
	db *gorm.DB
}

func (s StockRepositoryImpl) FindAllStocks() ([]dao.Stock, error) {
	return nil, nil
}

func (s StockRepositoryImpl) FindStockById(id int) (*dao.Stock, error) {
	return nil, nil
}

func (s StockRepositoryImpl) SaveStock(stock *dao.Stock) (*dao.Stock, error) {
	return nil, nil
}

func (s StockRepositoryImpl) DeleteStockById(id int) error {
	return nil
}

func StockRepositoryInit(db *gorm.DB) *StockRepositoryImpl {
	err := db.AutoMigrate(&dao.Stock{})
	if err != nil {
		log.Fatalf("Unable to migrate model changes %s", err)
	}
	return &StockRepositoryImpl{
		db: db,
	}
}
