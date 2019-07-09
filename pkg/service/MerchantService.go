package service

import (
	"fmt"

	"github.com/jinzhu/gorm"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload" // part of transport
)

// MerchantService interface
type MerchantService interface {
	GetMerchants() (*[]models.Merchant, error)
	ShowMerchant(string) (models.Merchant, error)
	CreateMerchant(payload.CreateMerchantRequest) payload.CreateMerchantResponse
	UpdateMerchant(payload.UpdateMerchantRequest, string) payload.UpdateMerchantResponse
	DeleteMerchant(string) payload.DeleteMerchantResponse
}

type merchantService struct{}

var query *gorm.DB

// NewMerchantService for
func NewMerchantService(db *gorm.DB) MerchantService {
	query = db
	return merchantService{}
}

//GetMerchants for get all merchant
func (merchantService) GetMerchants() (*[]models.Merchant, error) {
	var merchants []models.Merchant

	query.Find(&merchants)

	// defer query.Close()

	return &merchants, nil
}

//ShowMerchant show merchant by id
func (merchantService) ShowMerchant(id string) (models.Merchant, error) {
	var merchant models.Merchant
	query.Find(&merchant, id)
	return merchant, nil
}

func (merchantService) CreateMerchant(data payload.CreateMerchantRequest) payload.CreateMerchantResponse {
	merchant := models.Merchant{
		Subdomain:      data.Subdomain,
		BusinessName:   data.BusinessName,
		BusinessType:   data.BusinessType,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Email:          data.Email,
		Mobile:         data.Mobile,
		Password:       data.Password,
		Status:         data.Status,
		APISecretKey:   data.APISecretKey,
		Passcode:       data.Passcode,
		ProfilePicture: data.ProfilePicture,
		LocationID:     data.LocationID,
		IsStore:        data.IsStore,
		IsMitra:        data.IsMitra,
		ReferralCode:   data.ReferralCode,
		IsSupplier:     data.IsSupplier,
	}
	err := models.InsertMerchant(query, &merchant)
	fmt.Println(&merchant)
	if err != nil {
		return payload.CreateMerchantResponse{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	return payload.CreateMerchantResponse{
		Message:    "Merchant Successfully Created",
		StatusCode: 201,
	}
}

//UpdateMerchant for updated merchant
func (merchantService) UpdateMerchant(data payload.UpdateMerchantRequest, id string) payload.UpdateMerchantResponse {
	var merchant models.Merchant
	if query.Find(&merchant, id).RecordNotFound() {
		return payload.UpdateMerchantResponse{
			Message:    "Merchant Not Found",
			StatusCode: 404,
		}
	}
	merchant.Subdomain = data.Merchant.Subdomain
	merchant.BusinessName = data.Merchant.BusinessName
	merchant.BusinessType = data.Merchant.BusinessType
	merchant.FirstName = data.Merchant.FirstName
	merchant.LastName = data.Merchant.LastName
	merchant.Email = data.Merchant.Email
	merchant.Mobile = data.Merchant.Mobile
	merchant.Password = data.Merchant.Password
	merchant.Status = data.Merchant.Status
	merchant.APISecretKey = data.Merchant.APISecretKey
	merchant.Passcode = data.Merchant.Passcode
	merchant.ProfilePicture = data.Merchant.ProfilePicture
	merchant.LocationID = data.Merchant.LocationID
	merchant.IsStore = data.Merchant.IsStore
	merchant.IsMitra = data.Merchant.IsMitra
	merchant.ReferralCode = data.Merchant.ReferralCode
	merchant.IsSupplier = data.Merchant.IsSupplier

	query.Save(&merchant)

	return payload.UpdateMerchantResponse{
		Message:    "Merchant Successfully updated",
		StatusCode: 200,
	}
}

//DeleteMerchant delete data
func (merchantService) DeleteMerchant(id string) payload.DeleteMerchantResponse {
	var merchant models.Merchant
	fmt.Println("merchant", &merchant)
	fmt.Println("here", query)
	if query.Find(&merchant, id).RecordNotFound() {
		return payload.DeleteMerchantResponse{
			Message:    "Merchant Not Found",
			StatusCode: 404,
		}
	}
	query.Delete(&merchant)

	return payload.DeleteMerchantResponse{
		Message:    "Merchant Successfully Deleted",
		StatusCode: 200,
	}
}
