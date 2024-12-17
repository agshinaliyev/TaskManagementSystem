package model

import "gorm.io/gorm"

type Rate struct {
	ID       uint
	Currency string  `gorm:"type:varchar(3);unique_index"`
	Price    float64 `gorm:"type:decimal(10,2)"`
	gorm.Model
}
