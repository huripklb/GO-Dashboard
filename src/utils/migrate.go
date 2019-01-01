package utils

import (
	"github.com/delivery-api/model"
	"github.com/jinzhu/gorm"
)

// Migrate : migrate table for first initialization
// author : Huripto Sugandi
// created date : 7 Dec 2018
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.Apiuser{})
}
