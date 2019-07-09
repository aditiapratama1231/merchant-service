package endpoint

import "github.com/go-kit/kit/endpoint"

// Endpoints Register our endpoints here
type Endpoints struct {
	GetMerchantsEndpoint   endpoint.Endpoint
	ShowMerchantEndpoint   endpoint.Endpoint
	CreateMerchantEndpoint endpoint.Endpoint
	UpdateMerchantEndpoint endpoint.Endpoint
	DeleteMerchantEndpoint endpoint.Endpoint
}
