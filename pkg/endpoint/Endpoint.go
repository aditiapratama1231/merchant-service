package endpoint

import "github.com/go-kit/kit/endpoint"

// Register our endpoints here
type Endpoints struct {
	GetProductsEndpoint   endpoint.Endpoint
	ShowProductsEndpoint  endpoint.Endpoint
	CreateProductEndpoint endpoint.Endpoint
	UpdateProductEnpoint  endpoint.Endpoint
	DeleteProductEnpoint  endpoint.Endpoint

	GetMerchantsEndpoint   endpoint.Endpoint
	ShowMerchantEndpoint   endpoint.Endpoint
	CreateMerchantEndpoint endpoint.Endpoint
	UpdateMerchantEndpoint endpoint.Endpoint
	DeleteMerchantEndpoint endpoint.Endpoint
}
