package request

import (
	models "qasir-supplier/merchant/models"
)

type Merchant struct {
	Subdomain      string `json:"subdomain"`
	BusinessName   string `json:"business_name"`
	BusinessType   string `json:"business_type"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
	Mobile         string `json:"mobile"`
	Password       string `json:"password"`
	Status         string `json:"status"`
	APISecretKey   string `json:"api_secret_key"`
	Passcode       string `json:"passcode"`
	ProfilePicture string `json:"profile_picture"`
	LocationID     int64  `json:"location_id"`
	IsStore        int64  `json:"is_store"`
	IsMitra        int64  `json:"is_mitra"`
	ReferralCode   string `json:"referral_code"`
	IsSupplier     int64  `json:"is_supplier"`
}

//GetMerchantsRequest show all merchant
type GetMerchantsRequest struct{}

type GetMerchantsResponse struct {
	Merchants []models.Merchant `json:"merchants"`
	Err       string            `json:"eer,omitempty"`
}

//showMerchantRequest by id
type ShowMerchantRequest struct {
	Id string `json:"id"`
}

type ShowMerchantResponse struct {
	Merchant models.Merchant `json:"merchant,omitempty"`
	Err      string          `json:"err,omitempty"`
}

//CreateMerchantRequest create new merchant
type CreateMerchantRequest struct {
	Merchant
}

type CreateMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

// UpdateMerchantRequest update merchant
type UpdateMerchantRequest struct {
	Id string `json:"id"`
	Merchant
}

type UpdateMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//DeleteMerchantRequest delete merchant
type DeleteMerchantRequest struct {
	Id string `json:"id"`
}

type DeleteMerchantResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
