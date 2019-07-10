package service

import (
	"fmt"

	"github.com/jinzhu/gorm"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload" // part of transport
)

// OutletService interface
type OutletService interface {
	GetOutlets() (*[]models.Outlet, error)
	ShowOutlet(string) (models.Outlet, error)
	CreateOutlet(payload.CreateOutletRequest) payload.CreateOutletResponse
	UpdateOutlet(payload.UpdateOutletRequest, string) payload.UpdateOutletResponse
	DeleteOutlet(string) payload.DeleteOutletResponse
}

type outletService struct{}

// NewOutletService for
func NewOutletService(db *gorm.DB) OutletService {
	query = db
	return outletService{}
}

//GetOutlets for get all merchant
func (outletService) GetOutlets() (*[]models.Outlet, error) {
	var outlets []models.Outlet

	query.Find(&outlets)

	// defer query.Close()

	return &outlets, nil
}

//ShowMerchant show merchant by id
func (outletService) ShowOutlet(id string) (models.Outlet, error) {
	var outlet models.Outlet
	query.Find(&outlet, id)
	return outlet, nil
}

func (outletService) CreateOutlet(data payload.CreateOutletRequest) payload.CreateOutletResponse {
	outlet := models.Outlet{
		MerchantID:          data.MerchantID,
		Code:                data.Code,
		Name:                data.Name,
		Phone:               data.Phone,
		Image:               data.Image,
		PlaceAddress:        data.PlaceAddress,
		PlaceLongitude:      data.PlaceLongitude,
		PlaceLatitude:       data.PlaceLatitude,
		MapAddress:          data.MapAddress,
		PlaceLocation:       data.PlaceLocation,
		PlacePostalcode:     data.PlacePostalcode,
		ReceiptName:         data.ReceiptName,
		ReceiptShowPlace:    data.ReceiptShowPlace,
		ReceiptShowLocation: data.ReceiptShowLocation,
		ReceiptShowPhone:    data.ReceiptShowPhone,
		CreatedBy:           data.CreatedBy,
		UpdatedBy:           data.UpdatedBy,
		DeletedBy:           data.DeletedBy,
	}
	err := models.InsertOutlet(query, &outlet)
	fmt.Println(&outlet)
	if err != nil {
		return payload.CreateOutletResponse{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	return payload.CreateOutletResponse{
		Message:    "Outlet Successfully Created",
		StatusCode: 201,
	}
}

//UpdateOutlet for updated merchant
func (outletService) UpdateOutlet(data payload.UpdateOutletRequest, id string) payload.UpdateOutletResponse {
	var outlet models.Outlet
	if query.Find(&outlet, id).RecordNotFound() {
		return payload.UpdateOutletResponse{
			Message:    "Outlet Not Found",
			StatusCode: 404,
		}
	}
	outlet.MerchantID = data.Outlet.MerchantID
	outlet.Code = data.Outlet.Code
	outlet.Name = data.Outlet.Name
	outlet.Phone = data.Outlet.Phone
	outlet.Image = data.Outlet.Image
	outlet.PlaceAddress = data.Outlet.PlaceAddress
	outlet.PlaceLongitude = data.Outlet.PlaceLongitude
	outlet.PlaceLatitude = data.Outlet.PlaceLatitude
	outlet.MapAddress = data.Outlet.MapAddress
	outlet.PlaceLocation = data.Outlet.PlaceLocation
	outlet.PlacePostalcode = data.Outlet.PlacePostalcode
	outlet.ReceiptName = data.Outlet.ReceiptName
	outlet.ReceiptShowPlace = data.Outlet.ReceiptShowPlace
	outlet.ReceiptShowLocation = data.Outlet.ReceiptShowLocation
	outlet.ReceiptShowPhone = data.Outlet.ReceiptShowPhone
	outlet.CreatedBy = data.Outlet.CreatedBy
	outlet.UpdatedBy = data.Outlet.UpdatedBy
	outlet.DeletedBy = data.Outlet.DeletedBy

	query.Save(&outlet)

	return payload.UpdateOutletResponse{
		Message:    "Outlet Successfully updated",
		StatusCode: 200,
	}
}

//DeleteMerchant delete data
func (outletService) DeleteOutlet(id string) payload.DeleteOutletResponse {
	var outlet models.Outlet
	//fmt.Println("merchant", &outlet)
	//fmt.Println("here", query)
	if query.Find(&outlet, id).RecordNotFound() {
		return payload.DeleteOutletResponse{
			Message:    "Outlet Not Found",
			StatusCode: 404,
		}
	}
	query.Delete(&outlet)

	return payload.DeleteOutletResponse{
		Message:    "Outlet Successfully Deleted",
		StatusCode: 200,
	}
}