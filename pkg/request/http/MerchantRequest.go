package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	payload "qasir-supplier/merchant/pkg/request/payload"
)

//DecodeGetMerchantsRequest decode our incoming requests
func DecodeGetMerchantsRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.GetMerchantsRequest
	return req, nil
}

//DecodeShowMerchantRequest decode our incoming requests
func DecodeShowMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ShowMerchantRequest{
		Id: qs["id"],
	}
	return req, nil
}

// DecodeCreateMerchantRequest decode our incoming request
func DecodeCreateMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req payload.CreateMerchantRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

//DecodeUpdateMerchantRequest decode our incoming request
func DecodeUpdateMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.UpdateMerchantRequest{
		Id: qs["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeDeleteMerchantRequest decode our incoming request
func DecodeDeleteMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.DeleteMerchantRequest{
		Id: qs["id"],
	}
	return req, nil
}
