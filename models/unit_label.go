package models

import (
	"github.com/jinzhu/gorm"
)

// UnitLabel model
type UnitLabel struct {
	gorm.Model
	Name      string `gorm:"column:name" json:"name"`
	CreatedBy int64  `gorm:"column:created_by" json:"created_by"`
	UpdatedBy int64  `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy int64  `gorm:"column:deleted_by" json:"deleted_by"`
}
