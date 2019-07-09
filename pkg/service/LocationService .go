package service

import (
	"fmt"

	"github.com/jinzhu/gorm"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload" // part of transport
)

// LocationService interface
type LocationService interface {
	GetLocations() (*[]models.Location, error)
	ShowLocation(string) (models.Location, error)
	CreateLocation(payload.CreateLocationRequest) payload.CreateLocationResponse
	UpdateLocation(payload.UpdateLocationRequest, string) payload.UpdateLocationResponse
	DeleteLocation(string) payload.DeleteLocationResponse
}

type locationService struct{}

// NewLocationService for
func NewLocationService(db *gorm.DB) LocationService {
	query = db
	return locationService{}
}

//GetOutlets for get all merchant
func (locationService) GetLocations() (*[]models.Location, error) {
	var locations []models.Location

	query.Find(&locations)

	// defer query.Close()

	return &locations, nil
}

//ShowLocation show merchant by id
func (locationService) ShowLocation(id string) (models.Location, error) {
	var location models.Location
	query.Find(&location, id)
	return location, nil
}

func (locationService) CreateLocation(data payload.CreateLocationRequest) payload.CreateLocationResponse {
	location := models.Location{
		Name:       data.Name,
		Type:       data.Type,
		Parent:     data.Parent,
		PostalCode: data.PostalCode,
		MapName:    data.MapName,
		Image:      data.Image,
		Caption:    data.Caption,
	}
	err := models.InsertLocation(query, &location)
	fmt.Println(&location)
	if err != nil {
		return payload.CreateLocationResponse{
			Message:    err.Error(),
			StatusCode: 500,
		}
	}

	return payload.CreateLocationResponse{
		Message:    "Location Successfully Created",
		StatusCode: 201,
	}
}

//UpdateLocation for updated merchant
func (locationService) UpdateLocation(data payload.UpdateLocationRequest, id string) payload.UpdateLocationResponse {
	var location models.Location
	if query.Find(&location, id).RecordNotFound() {
		return payload.UpdateLocationResponse{
			Message:    "Location Not Found",
			StatusCode: 404,
		}
	}
	location.Name = data.Location.Name
	location.Type = data.Location.Type
	location.Parent = data.Location.Parent
	location.PostalCode = data.Location.PostalCode
	location.MapName = data.Location.MapName
	location.Image = data.Location.Image
	location.Caption = data.Location.Caption

	query.Save(&location)

	return payload.UpdateLocationResponse{
		Message:    "Location Successfully updated",
		StatusCode: 200,
	}
}

//DeleteLocation delete data
func (locationService) DeleteLocation(id string) payload.DeleteLocationResponse {
	var location models.Location
	//fmt.Println("merchant", &outlet)
	//fmt.Println("here", query)
	if query.Find(&location, id).RecordNotFound() {
		return payload.DeleteLocationResponse{
			Message:    "Location Not Found",
			StatusCode: 404,
		}
	}
	query.Delete(&location)

	return payload.DeleteLocationResponse{
		Message:    "Location Successfully Deleted",
		StatusCode: 200,
	}
}
