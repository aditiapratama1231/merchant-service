package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Outlet model
type Outlet struct {
	gorm.Model
	MerchantID          int64       `gorm:"column:merchant_id" json:"merchant_id"`
	Code                string      `gorm:"column:code" json:"code"`
	Name                string      `gorm:"column:name" json:"name"`
	Phone               string      `gorm:"column:phone" json:"phone"`
	Image               string      `gorm:"column:image" json:"image"`
	PlaceAddress        string      `gorm:"column:place_address" json:"place_address"`
	PlaceLongitude      string      `gorm:"column:place_longitude" json:"place_longitude"`
	PlaceLatitude       string      `gorm:"column:place_latitude" json:"place_latitude"`
	MapAddress          string      `gorm:"column:map_address" json:"map_address"`
	PlaceLocation       int64       `gorm:"column:place_location" json:"place_location"`
	PlacePostalcode     string      `gorm:"column:place_postalcode" json:"place_postalcode"`
	ReceiptName         string      `gorm:"column:receipt_name" json:"receipt_name"`
	ReceiptShowPlace    int8        `gorm:"column:receipt_show_place" json:"receipt_show_place"`
	ReceiptShowLocation int8        `gorm:"column:receipt_show_location" json:"receipt_show_location"`
	ReceiptShowPhone    int8        `gorm:"column:receipt_show_phone" json:"receipt_show_phone"`
	CreatedBy           int64       `gorm:"column:created_by" json:"created_by"`
	UpdatedBy           int64       `gorm:"column:updated_by" json:"updated_by"`
	DeletedBy           int64       `gorm:"column:deleted_by" json:"deleted_by"`
	Location            []*Location `gorm:"many2many:outlet_coverages;" json:"Location,omitempty"`
	Merchant            Merchant    `gorm:"foreignkey:merchant_id;association_foreignkey:id" json:"Merchant,omitempty"`
}

//InsertOutlet  for insert outlet to db
func InsertOutlet(db *gorm.DB, p *Outlet) (err error) {
	if err = db.Save(p).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
