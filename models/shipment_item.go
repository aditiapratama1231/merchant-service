package models

import (
	"github.com/jinzhu/gorm"
)

// ShipmentItem model
type ShipmentItem struct {
	gorm.Model
	ShipmentID  int64   `gorm:"column:shipment_id" json:"shipment_id"`
	OrderItemID int64   `gorm:"column:order_item_id" json:"order_item_id"`
	VariantID   int64   `gorm:"column:variant_id" json:"variant_id"`
	Price       float64 `gorm:"column:price" json:"price"`
	Qty         float64 `gorm:"column:qty" json:"qty"`
	Weight      float64 `gorm:"column:weight" json:"weight"`
	Name        string  `gorm:"column:name" json:"name"`
	Sku         string  `gorm:"column:sku" json:"sku"`
	Volume      float64 `gorm:"column:volume" json:"volume"`
}
