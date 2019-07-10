package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	payload "qasir-supplier/merchant/pkg/request/payload"
)

//DecodeGetLocationsRequest decode our incoming requests
func DecodeGetLocationsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetLocationsRequest
	return req, nil
}

//DecodeShowLocationRequest decode our incoming requests
func DecodeShowLocationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ShowLocationRequest{
		ID: qs["id"],
	}
	return req, nil
}

// DecodeCreateLocationRequest decode our incoming request
func DecodeCreateLocationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//DecodeUpdateLocationRequest decode our incoming request
func DecodeUpdateLocationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.UpdateLocationRequest{
		ID: qs["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeDeleteLocationRequest decode our incoming request
func DecodeDeleteLocationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.DeleteLocationRequest{
		ID: qs["id"],
	}
	return req, nil
}
