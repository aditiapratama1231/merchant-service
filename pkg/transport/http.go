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

	//Outlet handler
	getOutletsHandler := httptransport.NewServer(
		endpoints.GetOutletsEndpoint,
		request.DecodeGetOutletsRequest,
		request.EncodeResponse,
	)

	showOutletHandler := httptransport.NewServer(
		endpoints.ShowOutletEndpoint,
		request.DecodeShowOutletRequest,
		request.EncodeResponse,
	)

	createOutletHandler := httptransport.NewServer(
		endpoints.CreateOutletEndpoint,
		request.DecodeCreateOutletRequest,
		request.EncodeResponse,
	)

	updateOutletHandler := httptransport.NewServer(
		endpoints.UpdateOutletEndpoint,
		request.DecodeUpdateOutletRequest,
		request.EncodeResponse,
	)

	deleteOutletHandler := httptransport.NewServer(
		endpoints.DeleteOutletEndpoint,
		request.DecodeDeleteOutletRequest,
		request.EncodeResponse,
	)

	// merchant handler
	r.Handle("/merchants", getMerchantsHandler).Methods("GET")
	r.Handle("/merchants/create", createMerchantHandler).Methods("POST")
	r.Handle("/merchants/{id}", showMerchantHandler).Methods("GET")
	r.Handle("/merchants/{id}/update", updateMerchanthandler).Methods("PATCH")
	r.Handle("/merchants/{id}/delete", deleteMerchantHandler).Methods("DELETE")

	// outlet handler
	r.Handle("/outlets", getOutletsHandler).Methods("GET")
	r.Handle("/outlets/create", createOutletHandler).Methods("POST")
	r.Handle("/outlets/{id}", showOutletHandler).Methods("GET")
	r.Handle("/outlets/{id}/update", updateOutletHandler).Methods("PATCH")
	r.Handle("/outlets/{id}/delete", deleteOutletHandler).Methods("DELETE")

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
