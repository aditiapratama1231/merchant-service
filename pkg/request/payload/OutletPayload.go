package request

import (
	models "qasir-supplier/merchant/models"

	"github.com/jinzhu/gorm"
)

//Outlet struct
type Outlet struct {
	gorm.Model
	MerchantID          int64                   `json:"merchant_id"`
	Code                string                  `json:"code"`
	Name                string                  `json:"name"`
	Phone               string                  `json:"phone"`
	Image               string                  `json:"image"`
	PlaceAddress        string                  `json:"place_address"`
	PlaceLongitude      string                  `json:"place_longitude"`
	PlaceLatitude       string                  `json:"place_latitude"`
	MapAddress          string                  `json:"map_address"`
	PlaceLocation       int64                   `json:"place_location"`
	PlacePostalcode     string                  `json:"place_postalcode"`
	ReceiptName         string                  `json:"receipt_name"`
	ReceiptShowPlace    int8                    `json:"receipt_show_place"`
	ReceiptShowLocation int8                    `json:"receipt_show_location"`
	ReceiptShowPhone    int8                    `json:"receipt_show_phone"`
	CreatedBy           int64                   `json:"created_by"`
	UpdatedBy           int64                   `json:"updated_by"`
	DeletedBy           int64                   `json:"deleted_by"`
	OutletCoverages     []models.OutletCoverage `gorm:"foreignkey:outlet_id" json:"outlet_coverages,omitempty"`
}

//GetOutletsRequest show all outlet to our server
type GetOutletsRequest struct{}

//GetOutletsResponse response from our server
type GetOutletsResponse struct {
	Outlets []models.Outlet `json:"outlets"`
	Err     string          `json:"eer,omitempty"`
}

//ShowOutletRequest to our server
type ShowOutletRequest struct {
	ID string `json:"id"`
}

// type ShowOutletLocationsRequest struct {
// 	OutletID	string `json:`
// }

//ShowOutletResponse from our server
type ShowOutletResponse struct {
	Message string        `json:"message,omitempty"`
	Data    models.Outlet `json:"Data,omitempty"`
	Err     int32         `json:"err,omitempty"`
}

//ShowOutletLocationsResponse /
type ShowOutletLocationsResponse struct {
	Message string            `json:"message,omitempty"`
	Data    []models.Location `json:"Data,omitempty"`
	Err     int32             `json:"err,omitempty"`
}

//CreateOutletRequest to our server
type CreateOutletRequest struct {
	Outlet
}

//CreateOutletResponse from our server
type CreateOutletResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

// UpdateOutletRequest to our server
type UpdateOutletRequest struct {
	ID string `json:"id"`
	Outlet
}

// UpdateOutletResponse from our server
type UpdateOutletResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//DeleteOutletRequest to our server
type DeleteOutletRequest struct {
	ID string `json:"id"`
}

//DeleteOutletResponse from our server
type DeleteOutletResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}

//CreateOutletLocationRequest struct
type CreateOutletLocationRequest struct {
	OutletCoverage
}

//CreateOutletLocationResponse struct
type CreateOutletLocationResponse struct {
	Message    string `json:"message"`
	StatusCode int32  `json:"status_code"`
}
