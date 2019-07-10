package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Merchant model
type Merchant struct {
	gorm.Model
	Subdomain      string   `gorm:"column:subdomain" json:"subdomain"`
	BusinessName   string   `gorm:"column:business_name" json:"business_name"`
	BusinessType   string   `gorm:"column:business_type" json:"business_type"`
	FirstName      string   `gorm:"column:first_name" json:"first_name"`
	LastName       string   `gorm:"column:last_name" json:"last_name"`
	Email          string   `gorm:"column:email" json:"email"`
	Mobile         string   `gorm:"column:mobile" json:"mobile"`
	Password       string   `gorm:"column:password" json:"password"`
	Status         string   `gorm:"column:status" json:"status"`
	APISecretKey   string   `gorm:"column:api_secret_key" json:"api_secret_key"`
	Passcode       string   `gorm:"column:passcode" json:"passcode"`
	ProfilePicture string   `gorm:"column:profile_picture" json:"profile_picture"`
	LocationID     int64    `gorm:"column:location_id" json:"location_id"`
	IsStore        int64    `gorm:"column:is_store" json:"is_store"`
	IsMitra        int64    `gorm:"column:is_mitra" json:"is_mitra"`
	ReferralCode   string   `gorm:"column:referral_code" json:"referral_code"`
	IsSupplier     int64    `gorm:"column:is_supplier" json:"is_supplier"`
	Outlets        []Outlet `gorm:"foreignkey:merchant_id" json:"outlets"`
}

//InsertMerchant  for insert merchant to db
func InsertMerchant(db *gorm.DB, p *Merchant) (err error) {
	if err = db.Save(p).Error; err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
