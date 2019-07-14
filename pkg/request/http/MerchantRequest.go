package request

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	payload "qasir-supplier/merchant/pkg/request/payload"
)

//DecodeGetMerchantsCoverageRequest decode our incoming requests
func DecodeGetMerchantsCoverageRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	v := r.URL.Query()
	qs := mux.Vars(r)
	include := strings.Split(strings.Replace(v.Get("include"), " ", "", -1), ",")
	page, _ := strconv.ParseInt(v.Get("page"), 10, 32)
	limit, _ := strconv.ParseInt(v.Get("limit"), 10, 32)
	req := payload.GetMerchantsCoverageRequest{
		PaginationRequest: payload.PaginationRequest{
			Page:    uint32(page),
			Limit:   uint32(limit),
			Include: include,
		},
		LocationID: qs["location_id"],
		Type:       qs["type"],
	}
	return req, nil
}

//DecodeShowMerchantRequest decode our incoming requests
func DecodeShowMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.ShowMerchantRequest{
		ID: qs["id"],
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
		ID: qs["id"],
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// DecodeDeleteMerchantRequest decode our incoming request
func DecodeDeleteMerchantRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	qs := mux.Vars(r)
	req := payload.DeleteMerchantRequest{
		ID: qs["id"],
	}
	return req, nil
}
