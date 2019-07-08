package models

import (
	"github.com/jinzhu/gorm"
)

//OrderStatus model
type OrderStatus struct {
	gorm.Model
	Code  string `gorm:"column:code" json:"code"`
	Name  string `gorm:"column:name" json:"name"`
	Order int8   `gorm:"column:order" json:"order"`
}
