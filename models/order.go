package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order model
type Order struct {
	gorm.Model
	VariantID      int64     `gorm:"column:variant_id" json:"variant_id"`
	SupplierID     int64     `gorm:"column:supplier_id" json:"supplier_id"`
	MerchantID     int64     `gorm:"column:merchant_id" json:"merchant_id"`
	OutletID       int64     `gorm:"column:outletID" json:"outletID"`
	Subtotal       float64   `gorm:"column:subtotal" json:"subtotal"`
	DiscountAmount float64   `gorm:"column:discount_amount" json:"discount_amount"`
	GrandTotal     float64   `gorm:"column:grand_total" json:"grand_total"`
	OrderNo        string    `gorm:"column:order_no" json:"order_no"`
	Status         string    `gorm:"column:status" json:"status"`
	HasInvoice     int8      `gorm:"column:has_invoice" json:"has_invoice"`
	CanceledAt     time.Time `json:"canceled_at" db:"canceled_at"`
	ShippedAt      time.Time `json:"shipped_at" db:"shipped_at"`
	DeliveredAt    time.Time `json:"delivered_at" db:"delivered_at"`
	CompletedAt    time.Time `json:"completed_at" db:"completed_at"`
}
