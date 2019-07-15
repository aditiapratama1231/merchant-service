package endpoint

import "github.com/go-kit/kit/endpoint"

// Endpoints Register our endpoints here
type Endpoints struct {
	GetMerchantsCoverageEndpoint endpoint.Endpoint
	ShowMerchantEndpoint         endpoint.Endpoint
	CreateMerchantEndpoint       endpoint.Endpoint
	UpdateMerchantEndpoint       endpoint.Endpoint
	DeleteMerchantEndpoint       endpoint.Endpoint

	GetOutletsEndpoint           endpoint.Endpoint
	ShowOutletEndpoint           endpoint.Endpoint
	CreateOutletEndpoint         endpoint.Endpoint
	UpdateOutletEndpoint         endpoint.Endpoint
	DeleteOutletEndpoint         endpoint.Endpoint
	CreateOutletLocationEndpoint endpoint.Endpoint
	ShowOutletLocationsEndpoint  endpoint.Endpoint

	GetLocationsEndpoint        endpoint.Endpoint
	ShowLocationEndpoint        endpoint.Endpoint
	CreateLocationEndpoint      endpoint.Endpoint
	UpdateLocationEndpoint      endpoint.Endpoint
	DeleteLocationEndpoint      endpoint.Endpoint
	ShowLocationOutletsEndpoint endpoint.Endpoint
}
