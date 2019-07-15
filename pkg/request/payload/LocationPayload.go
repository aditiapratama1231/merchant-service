package request

import (
	models "qasir-supplier/merchant/models"
)

//OutletCoverage struct
type OutletCoverage struct {
	OutletID   int64 `json:"outlet_id"`
	LocationID int64 `json:"location_id"`
}

// Location struct
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

//GetLocationsResponse struct
type GetLocationsResponse struct {
	Locations []models.Location `json:"locations"`
	Err       string            `json:"eer,omitempty"`
}

//ShowLocationRequest by id
type ShowLocationRequest struct {
	ID string `json:"id"`
}

//ShowLocationResponse struct contain response from server
type ShowLocationResponse struct {
	Location models.Location `json:"location,omitempty"`
	Err      string          `json:"err,omitempty"`
}

//CreateLocationRequest create new merchant
type CreateLocationRequest struct {
	Location
}

//ShowLocationOutletsResponse /
type ShowLocationOutletsResponse struct {
	Message string          `json:"message,omitempty"`
	Data    []models.Outlet `json:"Data,omitempty"`
	Err     int32           `json:"err,omitempty"`
}

//CreateLocationResponse struct contain resp from server
type CreateLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

// UpdateLocationRequest update merchant
type UpdateLocationRequest struct {
	ID string `json:"id"`
	Location
}

//UpdateLocationResponse struct
type UpdateLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//DeleteLocationRequest delete merchant
type DeleteLocationRequest struct {
	ID string `json:"id"`
}

//DeleteLocationResponse struct
type DeleteLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
