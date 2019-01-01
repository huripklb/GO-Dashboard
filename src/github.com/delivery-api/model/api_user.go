package model

import (
	"github.com/jinzhu/gorm"
)

// Apiuser : table structure for Apiuser table, use gorm model
// author : Huripto Sugandi
// date created : 5 Dec 2018
type Apiuser struct {
	gorm.Model
	ID       int    `gorm:"auto_increment" gorm:"primary_key" json:"id"`
	Username string `gorm:"size:100" json:"username"`
	Password string `gorm:"size:50" json:"password"`
	Active   bool   `json:"active"`
}
