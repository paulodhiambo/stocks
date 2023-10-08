package repository

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"stocks/app/domain/dao"
)

type StockRepository interface {
	FindAllStocks(page, perPage int) (map[string]interface{}, error)
	FindStockById(id int) (*dao.Stock, error)
	SaveStock(stock *dao.Stock) (*dao.Stock, error)
	DeleteStockById(id int) error
}

type StockRepositoryImpl struct {
	db *gorm.DB
}

func (s StockRepositoryImpl) FindAllStocks(page, perPage int) (map[string]interface{}, error) {
	var stocks []dao.Stock
	db := s.db.Offset((page - 1) * perPage).Limit(perPage)
	if err := db.Find(&stocks).Error; err != nil {
		return nil, err
	}
	// Calculate the total number of records
	var totalRecords int64
	s.db.Model(&dao.Stock{}).Count(&totalRecords)
	response := map[string]interface{}{
		"stock":       stocks,
		"total":       totalRecords,
		"page":        page,
		"size":        perPage,
		"total_pages": (int(totalRecords) + perPage - 1) / perPage,
		"visible":     len(stocks),
	}
	return response, nil
}

func (s StockRepositoryImpl) FindStockById(id int) (*dao.Stock, error) {
	stock := dao.Stock{}
	if err := s.db.Where("trade_id = ?", id).First(&stock).Error; err != nil {
		log.Error("Got an error while finding stock by Id: ", err)
		return &dao.Stock{}, err
	}
	return &stock, nil
}

func (s StockRepositoryImpl) SaveStock(stock *dao.Stock) (*dao.Stock, error) {
	if err := s.db.Save(stock).Error; err != nil {
		log.Error("Got an error while trying to save a stock: ", err)
		return &dao.Stock{}, err
	}
	return stock, nil
}

func (s StockRepositoryImpl) DeleteStockById(id int) error {
	if err := s.db.Delete(&dao.Stock{}, id).Error; err != nil {
		log.Error("Got an error while trying to delete a user: ", err)
		return err
	}
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
