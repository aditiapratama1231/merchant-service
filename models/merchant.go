package models

import (
	"github.com/jinzhu/gorm"
)

// Merchant model
type Merchant struct {
	gorm.Model
	Subdomain      string `gorm:"column:subdomain" json:"subdomain"`
	BusinessName   string `gorm:"column:businnes_name" json:"businnes_name"`
	BusinessType   string `gorm:"column:businnes_type" json:"businnes_type"`
	FirstName      string `gorm:"column:first_name" json:"first_name"`
	LastName       string `gorm:"column:last_name" json:"last_name"`
	Email          string `gorm:"column:email" json:"email"`
	Mobile         string `gorm:"column:mobile" json:"mobile"`
	Password       string `gorm:"column:password" json:"password"`
	Status         string `gorm:"column:status" json:"status"`
	APISecretKey   string `gorm:"column:api_secret_key" json:"api_secret_key"`
	Passcode       string `gorm:"column:passcode" json:"passcode"`
	ProfilePicture string `gorm:"column:profile_picture" json:"profile_picture"`
	LocationID     int16  `gorm:"column:location_id" json:"location_id"`
	IsStore        int8   `gorm:"column:is_store" json:"is_store"`
	IsMitra        int8   `gorm:"column:is_mitra" json:"is_mitra"`
	ReferallCode   string `gorm:"column:referall_code" json:"referall_code"`
	IsSupplier     int8   `gorm:"column:is_supplier" json:"is_supplier"`
}
