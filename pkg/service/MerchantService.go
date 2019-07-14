package service

import (
	"fmt"

	"github.com/jinzhu/gorm"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload" // part of transport
)

// MerchantService interface
type MerchantService interface {
	GetMerchantsCoverage(payload.GetMerchantsCoverageRequest) (payload.GetMerchantsResponse, error)
	ShowMerchant(string) payload.ShowMerchantResponse
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
func (merchantService) GetMerchantsCoverage(data payload.GetMerchantsCoverageRequest) (payload.GetMerchantsResponse, error) {
	var (
		outlets  []models.Outlet
		location models.Location
	)
	//pagination
	offset := (data.PaginationRequest.Page - 1) * data.PaginationRequest.Limit

	//1:province,2:city,3:district,4:subdistrict
	if prd := query.Where("id=? AND type = ? ", data.LocationID, data.Type).First(&location); prd.Error != nil {
		return payload.GetMerchantsResponse{
			Message: "Location Not Available",
			Err:     404,
		}, nil
	}
	//cannot add count how data is get
	query.Model(location).Association("Outlet").Find(&outlets)

	//query.Limit(data.PaginationRequest.Limit).Offset(offset).Find(&outlets)
	for i := range outlets {
		query.Model(outlets[i]).Related(&outlets[i].Merchant)
	}

	return payload.GetMerchantsResponse{
		CurrentPage: data.PaginationRequest.Page,
		Data:        outlets,
		From:        offset + 1,
		To:          data.PaginationRequest.Page * data.PaginationRequest.Limit,
		PerPage:     data.PaginationRequest.Limit,
		//Total:       count,
	}, nil

}

//ShowMerchant show merchant by id
func (merchantService) ShowMerchant(id string) payload.ShowMerchantResponse {
	var (
		merchant payload.Merchant
	)

	if prd := query.Find(&merchant, id); prd.Error != nil {
		return payload.ShowMerchantResponse{
			Message: "Merchant Not Found",
			Err:     true,
		}
	}
	query.Model(merchant).Related(&merchant.Outlets)

	return payload.ShowMerchantResponse{
		Message: "Merchant Successfully Loaded",
		Data:    merchant,
		Err:     false,
	}
}

func (merchantService) CreateMerchant(data payload.CreateMerchantRequest) payload.CreateMerchantResponse {
	//field status ('CREATED', 'ACTIVE', 'INACTIVE', 'SUSPENDED')

	if data.Status != "CREATED" && data.Status != "ACTIVE" && data.Status != "INACTIVE" && data.Status != "SUSPENDED" {
		return payload.CreateMerchantResponse{
			Message:    "Field Status no valid",
			StatusCode: 422,
		}
	}

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
	fmt.Println("status", data.Merchant.Status)
	if data.Merchant.Status != "CREATED" && data.Merchant.Status != "ACTIVE" && data.Merchant.Status != "INACTIVE" && data.Merchant.Status != "SUSPENDED" {
		return payload.UpdateMerchantResponse{
			Message:    "Field Status no valid",
			StatusCode: 422,
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
