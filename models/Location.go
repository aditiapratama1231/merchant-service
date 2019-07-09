package models

import (
	"github.com/jinzhu/gorm"
)

// Location model
type Location struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	Type       int16  `gorm:"column:type" json:"type"`
	Parent     int64  `gorm:"column:parent" json:"parent"`
	PostalCode string `gorm:"column:postal_code" json:"postal_code"`
	MapName    string `gorm:"column:map_name" json:"map_name"`
	Image      string `gorm:"column:image" json:"image"`
	Caption    string `gorm:"column:caption" json:"caption"`
}