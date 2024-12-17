package db

import "gorm.io/gorm"

type Repository interface {
	FindRates()
	FindRatesByCurrency(currency string)
	UpdateRatesByCurrency(currency string, price float64)
}

func NewRepository(db *gorm.DB) Repository {
	return repository{db: db}

}

type repository struct {
	db *gorm.DB
}

func (r repository) FindRates() {

}
func (r repository) FindRatesByCurrency(currency string) {

}
func (r repository) UpdateRatesByCurrency(currency string, price float64) {

}
