package models

import (
	"github.com/jinzhu/gorm"
)

// InvoiceItem model
type InvoiceItem struct {
	gorm.Model
	InvoiceID       int64   `gorm:"column:invoice_id" json:"invoice_id"`
	VariantID       int64   `gorm:"column:variant_id" json:"variant_id"`
	Qty             float64 `gorm:"column:qty" json:"qty"`
	Price           float64 `gorm:"column:price" json:"price"`
	Sku             string  `gorm:"column:sku" json:"sku"`
	Name            string  `gorm:"column:name" json:"name"`
	InvoiceItemscol string  `gorm:"column:privileges" json:"privileges"`
}

// // Insert Product To Database
// func InsertProduct(db *gorm.DB, p *Product) (err error) {
// 	if err = db.Save(p).Error; err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	return nil
// }
