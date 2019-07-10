package endpoint

import (
	"context"

	payload "qasir-supplier/merchant/pkg/request/payload"
	"qasir-supplier/merchant/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

//MakeGetMerchantsEndpoint make endpoint
func MakeGetMerchantsEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.GetMerchantsRequest)
		d, err := srv.GetMerchants(req)
		if err != nil {
			return d, nil
		}
		return d, nil
	}
}

func MakeShowMerchantEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowMerchantRequest)
		//fmt.Println("here")
		d := srv.ShowMerchant(req.ID)
		return d, nil

	}
}

func MakeCreateMerchantEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateMerchantRequest)
		d := srv.CreateMerchant(req)
		return payload.CreateMerchantResponse{d.Message, d.StatusCode}, nil
	}
}

func MakeUpdateMerchantEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.UpdateMerchantRequest)
		d := srv.UpdateMerchant(req, req.ID)
		return payload.UpdateMerchantResponse{d.Message, d.StatusCode}, nil
	}
}

func MakeDeleteMerchantEndpoint(srv service.MerchantService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.DeleteMerchantRequest)
		d := srv.DeleteMerchant(req.ID)
		return payload.DeleteMerchantResponse{d.Message, d.StatusCode}, nil
	}
}

// func (e Endpoints) GetMerchants(ctx context.Context) ([]models.Merchant, error) {
// 	req := payload.GetMerchantsRequest{}
// 	resp, err := e.GetMerchantsEndpoint(ctx, req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	getResp := resp.(payload.GetMerchantsResponse)
// 	if getResp.Err != "" {
// 		return nil, errors.New(getResp.Err)
// 	}
// 	return getResp.Merchants, nil
// }

func (e Endpoints) ShowMerchant(ctx context.Context) payload.ShowMerchantResponse {
	req := payload.ShowMerchantRequest{}
	resp, _ := e.ShowMerchantEndpoint(ctx, req)
	getResp := resp.(payload.ShowMerchantResponse)
	return getResp
}

func (e Endpoints) CreateMerchant(ctx context.Context, data payload.CreateMerchantRequest) (payload.CreateMerchantResponse, error) {
	req := payload.CreateMerchantRequest{}
	resp, err := e.CreateMerchantEndpoint(ctx, req)
	if err != nil {
		return payload.CreateMerchantResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.CreateMerchantResponse)
	return getResp, nil
}

func (e Endpoints) UpdateMerchant(ctx context.Context, data payload.UpdateMerchantResponse) (payload.UpdateMerchantResponse, error) {
	req := payload.UpdateMerchantRequest{}
	resp, err := e.UpdateMerchantEndpoint(ctx, req)
	if err != nil {
		return payload.UpdateMerchantResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.UpdateMerchantResponse)
	return getResp, nil
}

func (e Endpoints) DeleteMerchant(ctx context.Context) (payload.DeleteMerchantResponse, error) {
	req := payload.DeleteMerchantRequest{}
	resp, err := e.DeleteMerchantEndpoint(ctx, req)
	if err != nil {
		return payload.DeleteMerchantResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.DeleteMerchantResponse)
	return getResp, nil
}
