package models

import "gorm.io/gorm"

type Stock struct {
	gorm.Model
	Code  string  `gorm:"uniqueIndex" json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
