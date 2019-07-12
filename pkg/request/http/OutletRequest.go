package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	payload "qasir-supplier/merchant/pkg/request/payload"
)

//DecodeGetOutletsRequest decode our incoming requests
func DecodeGetOutletsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetOutletsRequest
	return req, nil
}

//DecodeShowOutletRequest decode our incoming requests
func DecodeShowOutletRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ShowOutletRequest{
		ID: qs["id"],
	}
	return req, nil
}

// DecodeCreateOutletRequest decode our incoming request
func DecodeCreateOutletRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateOutletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//DecodeUpdateOutletRequest decode our incoming request
func DecodeUpdateOutletRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.UpdateOutletRequest{
		ID: qs["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeDeleteOutletRequest decode our incoming request
func DecodeDeleteOutletRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.DeleteOutletRequest{
		ID: qs["id"],
	}
	return req, nil
}

// DecodeCreateOutletLocationRequest decode our incoming request
func DecodeCreateOutletLocationRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateOutletLocationRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
