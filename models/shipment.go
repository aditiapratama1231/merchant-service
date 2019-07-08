package models

import (
	"github.com/jinzhu/gorm"
)

// Shipment model
type Shipment struct {
	gorm.Model
	OrderID      int64   `gorm:"column:order_id" json:"order_id"`
	OutletID     int64   `gorm:"column:outlet_id" json:"outlet_id"`
	TotalWeight  float64 `gorm:"column:total_weight" json:"total_weight"`
	TotalQty     float64 `gorm:"column:total_qty" json:"total_qty"`
	ShippingNo   string  `gorm:"column:shipping_no" json:"shipping_no"`
	ShippingNote string  `gorm:"column:shipping_note" json:"shipping_note"`
	CustomerNote string  `gorm:"column:customer_note" json:"customer_note"`
	Status       int8    `gorm:"column:status" json:"status"`
}
