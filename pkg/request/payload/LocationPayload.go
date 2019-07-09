package request

import (
	models "qasir-supplier/merchant/models"
)

type OutletCoverage struct {
	OutletID   int64 `json:"outlet_id"`
	LocationID int64 `json:"location_id"`
}

type Location struct {
	Name       string `json:"name"`
	Type       int16  `json:"type"`
	Parent     int64  `json:"parent"`
	PostalCode string `json:"postal_code"`
	MapName    string `json:"map_name"`
	Image      string `json:"image"`
	Caption    string `json:"caption"`
}

//GetLocationsRequest show all merchant
type GetLocationsRequest struct{}

type GetLocationsResponse struct {
	Locations []models.Location `json:"location"`
	Err       string            `json:"eer,omitempty"`
}

//ShowLocationRequest by id
type ShowLocationRequest struct {
	Id string `json:"id"`
}

type ShowLocationResponse struct {
	Location models.Location `json:"location,omitempty"`
	Err      string          `json:"err,omitempty"`
}

//CreateLocationRequest create new merchant
type CreateLocationRequest struct {
	Location
}

type CreateLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

// UpdateLocationRequest update merchant
type UpdateLocationRequest struct {
	Id string `json:"id"`
	Location
}

type UpdateLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//DeleteLocationRequest delete merchant
type DeleteLocationRequest struct {
	Id string `json:"id"`
}

type DeleteLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
