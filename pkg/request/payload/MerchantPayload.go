package request

import (
	models "qasir-supplier/merchant/models"
)

//Merchant struct contain variable
type Merchant struct {
	ID             uint            `json:"ID"`
	Subdomain      string          `json:"subdomain"`
	BusinessName   string          `json:"business_name"`
	BusinessType   string          `json:"business_type"`
	FirstName      string          `json:"first_name"`
	LastName       string          `json:"last_name"`
	Email          string          `json:"email"`
	Mobile         string          `json:"mobile"`
	Password       string          `json:"password"`
	Status         string          `json:"status"`
	APISecretKey   string          `json:"api_secret_key"`
	Passcode       string          `json:"passcode"`
	ProfilePicture string          `json:"profile_picture"`
	LocationID     int64           `json:"location_id"`
	IsStore        int64           `json:"is_store"`
	IsMitra        int64           `json:"is_mitra"`
	ReferralCode   string          `json:"referral_code"`
	IsSupplier     int64           `json:"is_supplier"`
	Outlets        []models.Outlet `gorm:"foreignkey:merchant_id" json:"outlets"`
}

//GetMerchantsRequest show all merchant
type GetMerchantsRequest struct {
	PaginationRequest
}

// GetMerchantsResponse struct contain response from server
type GetMerchantsResponse struct {
	Err         string            `json:"err,omitempty"`
	CurrentPage uint32            `json:"current_page"`
	From        uint32            `json:"from"`
	PerPage     uint32            `json:"per_page"`
	Total       uint32            `json:"total"`
	To          uint32            `json:"to"`
	Data        []models.Merchant `json:"data"`
}

//MerchantRequest struct
type MerchantRequest struct {
	Merchant       Merchant       `json:"product,omitempty"`
	Outlet         Outlet         `json:"outlet,omitempty"`
	OutletCoverage OutletCoverage `json:"outlet_coverage,omitempty"`
	Location       Location       `json:"location,omitempty"`
}

//ShowMerchantRequest by id
type ShowMerchantRequest struct {
	ID string `json:"id"`
}

//ShowMerchantResponse contain variable response from server
type ShowMerchantResponse struct {
	Message string   `json:"message"`
	Data    Merchant `json:"data"`
	Err     bool     `json:"err"`
}

//CreateMerchantRequest create new merchant
type CreateMerchantRequest struct {
	Merchant
}

//CreateMerchantResponse contain response from server
type CreateMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

// UpdateMerchantRequest update merchant
type UpdateMerchantRequest struct {
	ID string `json:"id"`
	Merchant
}

// UpdateMerchantResponse contain response from server
type UpdateMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//DeleteMerchantRequest delete merchant
type DeleteMerchantRequest struct {
	ID string `json:"id"`
}

// DeleteMerchantResponse contain response from server
type DeleteMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
