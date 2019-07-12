package endpoint

import (
	"context"
	"errors"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload"
	"qasir-supplier/merchant/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

//MakeGetOutletsEndpoint make endpoint
func MakeGetOutletsEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, err := srv.GetOutlets()
		if err != nil {
			return payload.GetOutletsResponse{nil, err.Error()}, nil
		}
		return payload.GetOutletsResponse{*d, ""}, nil
	}
}

//MakeShowOutletEndpoint make
func MakeShowOutletEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowOutletRequest)
		//fmt.Println("here")
		d, err := srv.ShowOutlet(req.ID)

		if err != nil {
			return payload.ShowOutletResponse{d, err.Error()}, nil
		}
		return payload.ShowOutletResponse{d, ""}, nil
	}
}

//MakeCreateOutletEndpoint //
func MakeCreateOutletEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateOutletRequest)
		d := srv.CreateOutlet(req)
		return payload.CreateOutletResponse{d.Message, d.StatusCode}, nil
	}
}

//MakeUpdateOutletEndpoint //
func MakeUpdateOutletEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.UpdateOutletRequest)
		d := srv.UpdateOutlet(req, req.ID)
		return payload.UpdateOutletResponse{d.Message, d.StatusCode}, nil
	}
}

// MakeDeleteOutletEndpoint //
func MakeDeleteOutletEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.DeleteOutletRequest)
		d := srv.DeleteOutlet(req.ID)
		return payload.DeleteOutletResponse{d.Message, d.StatusCode}, nil
	}
}

//MakeCreateOutletLocationEndpoint /
func MakeCreateOutletLocationEndpoint(srv service.OutletService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateOutletLocationRequest)
		d := srv.CreateOutletLocation(req)
		return payload.CreateOutletLocationResponse{d.Message, d.StatusCode}, nil
	}
}

//GetOutlets //
func (e Endpoints) GetOutlets(ctx context.Context) ([]models.Outlet, error) {
	req := payload.GetOutletsRequest{}
	resp, err := e.GetOutletsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(payload.GetOutletsResponse)
	if getResp.Err != "" {
		return nil, errors.New(getResp.Err)
	}
	return getResp.Outlets, nil
}

// ShowOutlet //
func (e Endpoints) ShowOutlet(ctx context.Context) (models.Outlet, error) {
	req := payload.ShowOutletRequest{}
	resp, _ := e.ShowOutletEndpoint(ctx, req)
	getResp := resp.(payload.ShowOutletResponse)
	return getResp.Outlet, nil
}

//CreateOutlet //
func (e Endpoints) CreateOutlet(ctx context.Context, data payload.CreateOutletRequest) (payload.CreateOutletResponse, error) {
	req := payload.CreateOutletRequest{}
	resp, err := e.CreateOutletEndpoint(ctx, req)
	if err != nil {
		return payload.CreateOutletResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.CreateOutletResponse)
	return getResp, nil
}

//UpdateOutlet //
func (e Endpoints) UpdateOutlet(ctx context.Context, data payload.UpdateOutletResponse) (payload.UpdateOutletResponse, error) {
	req := payload.UpdateOutletRequest{}
	resp, err := e.UpdateOutletEndpoint(ctx, req)
	if err != nil {
		return payload.UpdateOutletResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.UpdateOutletResponse)
	return getResp, nil
}

//DeleteOutlet //
func (e Endpoints) DeleteOutlet(ctx context.Context) (payload.DeleteOutletResponse, error) {
	req := payload.DeleteOutletRequest{}
	resp, err := e.DeleteOutletEndpoint(ctx, req)
	if err != nil {
		return payload.DeleteOutletResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.DeleteOutletResponse)
	return getResp, nil
}

//CreateOutletLocation //
func (e Endpoints) CreateOutletLocation(ctx context.Context, data payload.CreateOutletLocationRequest) (payload.CreateOutletLocationResponse, error) {
	req := payload.CreateOutletLocationRequest{}
	resp, err := e.CreateOutletEndpoint(ctx, req)
	if err != nil {
		return payload.CreateOutletLocationResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.CreateOutletLocationResponse)
	return getResp, nil
}
