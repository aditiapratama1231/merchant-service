package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// OutletCoverage model
type OutletCoverage struct {
	gorm.Model
	OutletID   int64    `gorm:"column:outlet_id" json:"outlet_id"`
	LocationID int64    `gorm:"column:location_id" json:"location_id"`
	Location   Location `gorm:"foreignkey:location_id" json:"location"`
}

// InsertOutletCoverage /
func InsertOutletCoverage(db *gorm.DB, p *OutletCoverage) (err error) {
	if err = db.Save(p).Error; err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
