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
	getMerchantsCoverageHandler := httptransport.NewServer(
		endpoints.GetMerchantsCoverageEndpoint,
		request.DecodeGetMerchantsCoverageRequest,
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

	showOutletLocationsHandler := httptransport.NewServer(
		endpoints.ShowOutletLocationsEndpoint,
		request.DecodeShowOutletLocationsRequest,
		request.EncodeResponse,
	)

	createOutletHandler := httptransport.NewServer(
		endpoints.CreateOutletEndpoint,
		request.DecodeCreateOutletRequest,
		request.EncodeResponse,
	)

	createOutletLocationHandler := httptransport.NewServer(
		endpoints.CreateOutletLocationEndpoint,
		request.DecodeCreateOutletLocationRequest,
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

	showLocationOutletsHandler := httptransport.NewServer(
		endpoints.ShowLocationOutletsEndpoint,
		request.DecodeShowLocationRequestOutlets,
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
	s.Handle("/{location_id}/{type}/locations", getMerchantsCoverageHandler).Methods("GET")
	s.Handle("/create", createMerchantHandler).Methods("POST")
	s.Handle("/{id}", showMerchantHandler).Methods("GET")
	s.Handle("/{id}/update", updateMerchanthandler).Methods("PATCH")
	s.Handle("/{id}/delete", deleteMerchantHandler).Methods("DELETE")

	// outlet handler
	s.Handle("/outlets/", getOutletsHandler).Methods("GET")
	s.Handle("/outlets/create", createOutletHandler).Methods("POST")
	s.Handle("/outlets/{id}", showOutletHandler).Methods("GET")
	s.Handle("/outlets/{id}/update", updateOutletHandler).Methods("PATCH")
	s.Handle("/outlets/{id}/delete", deleteOutletHandler).Methods("DELETE")
	//handle outlet can add location, payload outlet id and location id post send to outlet_coverage
	s.Handle("/outlets/add/locations", createOutletLocationHandler).Methods("POST")
	// //handle show outlet  with outlet id ? and its location
	s.Handle("/outlets/{id}/locations/", showOutletLocationsHandler).Methods("GET")

	// location handler
	s.Handle("/locations/", getLocationsHandler).Methods("GET")
	s.Handle("/locations/create", createLocationHandler).Methods("POST")
	s.Handle("/locations/{id}", showLocationHandler).Methods("GET")
	s.Handle("/locations/{id}/update", updateLocationHandler).Methods("PATCH")
	s.Handle("/locations/{id}/delete", deleteLocationHandler).Methods("DELETE")

	// // handle show location with location_id ? and its location
	s.Handle("/locations/{id}/outlets", showLocationOutletsHandler).Methods("GET")

	return s
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
