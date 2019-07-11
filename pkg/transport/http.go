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
	s := r.PathPrefix("/merchants").Subrouter()
	s.Use(commonMiddleware) // @see https://stackoverflow.com/a/51456342

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

	//Location handler
	getLocationsHandler := httptransport.NewServer(
		endpoints.GetLocationsEndpoint,
		request.DecodeGetLocationsRequest,
		request.EncodeResponse,
	)

	showLocationHandler := httptransport.NewServer(
		endpoints.ShowLocationEndpoint,
		request.DecodeShowLocationRequest,
		request.EncodeResponse,
	)

	createLocationHandler := httptransport.NewServer(
		endpoints.CreateLocationEndpoint,
		request.DecodeCreateLocationRequest,
		request.EncodeResponse,
	)

	updateLocationHandler := httptransport.NewServer(
		endpoints.UpdateLocationEndpoint,
		request.DecodeUpdateLocationRequest,
		request.EncodeResponse,
	)

	deleteLocationHandler := httptransport.NewServer(
		endpoints.DeleteLocationEndpoint,
		request.DecodeDeleteLocationRequest,
		request.EncodeResponse,
	)

	// merchant handler
	s.Handle("", getMerchantsHandler).Methods("GET")
	s.Handle("/create", createMerchantHandler).Methods("POST")
	s.Handle("/{id}", showMerchantHandler).Methods("GET")
	s.Handle("/{id}/update", updateMerchanthandler).Methods("PATCH")
	s.Handle("/{id}/delete", deleteMerchantHandler).Methods("DELETE")

	// outlet handler
	s.Handle("/outlets/", getOutletsHandler).Methods("GET")
	s.Handle("/outlet/create", createOutletHandler).Methods("POST")
	s.Handle("/outlet/{id}", showOutletHandler).Methods("GET")
	s.Handle("/outlet/{id}/update", updateOutletHandler).Methods("PATCH")
	s.Handle("/outlet/{id}/delete", deleteOutletHandler).Methods("DELETE")

	// location handler
	s.Handle("/locations/", getLocationsHandler).Methods("GET")
	s.Handle("/location/create", createLocationHandler).Methods("POST")
	s.Handle("/location/{id}", showLocationHandler).Methods("GET")
	s.Handle("/location/{id}/update", updateLocationHandler).Methods("PATCH")
	s.Handle("/location/{id}/delete", deleteLocationHandler).Methods("DELETE")

	return s
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
