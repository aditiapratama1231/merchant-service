package transport

import (
	"context"
	"net/http"

	"qasir-supplier/merchant/pkg/endpoint"
	request "qasir-supplier/merchant/pkg/request/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

	// merchant handler
	getMerchantsHandler := httptransport.NewServer(
		endpoints.GetMerchantsEndpoint,
		request.DecodeGetMerchantsRequest,
		request.EncodeResponse,
	)

	showMerchantHandler := httptransport.NewServer(
		endpoints.ShowMerchantEndpoint,
		request.DecodeShowMerchantRequest,
		request.EncodeResponse,
	)

	createMerchantHandler := httptransport.NewServer(
		endpoints.CreateMerchantEndpoint,
		request.DecodeCreateMerchantRequest,
		request.EncodeResponse,
	)

	updateMerchanthandler := httptransport.NewServer(
		endpoints.UpdateMerchantEndpoint,
		request.DecodeUpdateMerchantRequest,
		request.EncodeResponse,
	)

	deleteMerchantHandler := httptransport.NewServer(
		endpoints.DeleteMerchantEndpoint,
		request.DecodeDeleteMerchantRequest,
		request.EncodeResponse,
	)

	// products handlers
	getProductsHandler := httptransport.NewServer(
		endpoints.GetProductsEndpoint,
		request.DecodeGetProductsRequest,
		request.EncodeResponse,
	)

	showProductHandler := httptransport.NewServer(
		endpoints.ShowProductsEndpoint,
		request.DecodeShowProductRequest,
		request.EncodeResponse,
	)

	createProductHandler := httptransport.NewServer(
		endpoints.CreateProductEndpoint,
		request.DecodeCreateProductRequest,
		request.EncodeResponse,
	)

	updateProductHandler := httptransport.NewServer(
		endpoints.UpdateProductEnpoint,
		request.DecodeUpdateProductRequest,
		request.EncodeResponse,
	)

	deleteProductHandler := httptransport.NewServer(
		endpoints.DeleteProductEnpoint,
		request.DecodeDeleteProductRequest,
		request.EncodeResponse,
	)

	// products client endpoints
	r.Handle("/products", getProductsHandler).Methods("GET")
	r.Handle("/products/create", createProductHandler).Methods("POST")
	r.Handle("/products/{id}", showProductHandler).Methods("GET")
	r.Handle("/products/{id}/update", updateProductHandler).Methods("PATCH")
	r.Handle("/products/{id}/delete", deleteProductHandler).Methods("DELETE")

	// merchant handler
	r.Handle("/merchants", getMerchantsHandler).Methods("GET")
	r.Handle("/merchants/create", createMerchantHandler).Methods("POST")
	r.Handle("/merchants/{id}", showMerchantHandler).Methods("GET")
	r.Handle("/merchants/{id}/update", updateMerchanthandler).Methods("PATCH")
	r.Handle("/merchants/{id}/delete", deleteMerchantHandler).Methods("DELETE")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
