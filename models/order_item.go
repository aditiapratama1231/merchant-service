package models

import (
	"github.com/jinzhu/gorm"
)

// OrderItem model
type OrderItem struct {
	gorm.Model
	VariantID     int64   `gorm:"column:variant_id" json:"variant_id"`
	OrderID       int64   `gorm:"column:order_id" json:"order_id"`
	Qty           float64 `gorm:"column:qty" json:"qty"`
	QtyInvoice    float64 `gorm:"column:qty_invoice" json:"qty_invoice"`
	QtyShipped    float64 `gorm:"column:qty_shipped" json:"qty_shipped"`
	QtyCanceled   float64 `gorm:"column:qty_canceled" json:"qty_canceled"`
	QtyRefunded   float64 `gorm:"column:qty_refunded" json:"qty_refunded"`
	PriceSell     float64 `gorm:"column:price_sell" json:"price_sell"`
	PriceSellUnit float64 `gorm:"column:price_sell_unit" json:"price_sellUnit"`
	PriceBase     float64 `gorm:"column:price_base" json:"price_base"`
	PriceBaseUnit float64 `gorm:"column:price_base_unit" json:"price_base_unit"`
	Sku           string  `gorm:"column:sku" json:"sku"`
	Name          string  `gorm:"column:name" json:"name"`
}
