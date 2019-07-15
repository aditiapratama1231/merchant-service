package endpoint

import (
	"context"
	"errors"

	models "qasir-supplier/merchant/models"
	payload "qasir-supplier/merchant/pkg/request/payload"
	"qasir-supplier/merchant/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

//MakeGetLocationsEndpoint make endpoint
func MakeGetLocationsEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		d, err := srv.GetLocations()
		if err != nil {
			return payload.GetLocationsResponse{nil, err.Error()}, nil
		}
		return payload.GetLocationsResponse{*d, ""}, nil
	}
}

//MakeShowLocationEndpoint make
func MakeShowLocationEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowLocationRequest)
		//fmt.Println("here")
		d, err := srv.ShowLocation(req.ID)

		if err != nil {
			return payload.ShowLocationResponse{d, err.Error()}, nil
		}
		return payload.ShowLocationResponse{d, ""}, nil
	}
}

//MakeShowLocationsOutletEndpoint /
func MakeShowLocationsOutletsEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.ShowLocationRequest)
		d, err := srv.ShowLocationOutlets(req.ID)
		if err != nil {
			return d, nil
		}
		return d, nil
	}
}

//MakeCreateLocationEndpoint //
func MakeCreateLocationEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.CreateLocationRequest)
		d := srv.CreateLocation(req)
		return payload.CreateLocationResponse{d.Message, d.StatusCode}, nil
	}
}

//MakeUpdateLocationEndpoint //
func MakeUpdateLocationEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.UpdateLocationRequest)
		d := srv.UpdateLocation(req, req.ID)
		return payload.UpdateLocationResponse{d.Message, d.StatusCode}, nil
	}
}

//MakeDeleteLocationEndpoint //
func MakeDeleteLocationEndpoint(srv service.LocationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(payload.DeleteLocationRequest)
		d := srv.DeleteLocation(req.ID)
		return payload.DeleteLocationResponse{d.Message, d.StatusCode}, nil
	}
}

//GetLocations //
func (e Endpoints) GetLocations(ctx context.Context) ([]models.Location, error) {
	req := payload.GetLocationsRequest{}
	resp, err := e.GetLocationsEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}
	getResp := resp.(payload.GetLocationsResponse)
	if getResp.Err != "" {
		return nil, errors.New(getResp.Err)
	}
	return getResp.Locations, nil
}

// ShowLocation //
func (e Endpoints) ShowLocation(ctx context.Context) (models.Location, error) {
	req := payload.ShowLocationRequest{}
	resp, _ := e.ShowLocationEndpoint(ctx, req)
	getResp := resp.(payload.ShowLocationResponse)
	return getResp.Location, nil
}

// CreateLocation //
func (e Endpoints) CreateLocation(ctx context.Context, data payload.CreateLocationRequest) (payload.CreateLocationResponse, error) {
	req := payload.CreateLocationRequest{}
	resp, err := e.CreateLocationEndpoint(ctx, req)
	if err != nil {
		return payload.CreateLocationResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.CreateLocationResponse)
	return getResp, nil
}

// UpdateLocation //
func (e Endpoints) UpdateLocation(ctx context.Context, data payload.UpdateLocationRequest) (payload.UpdateLocationResponse, error) {
	req := payload.UpdateLocationRequest{}
	resp, err := e.UpdateLocationEndpoint(ctx, req)
	if err != nil {
		return payload.UpdateLocationResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.UpdateLocationResponse)
	return getResp, nil
}

//DeleteLocation //
func (e Endpoints) DeleteLocation(ctx context.Context) (payload.DeleteLocationResponse, error) {
	req := payload.DeleteLocationRequest{}
	resp, err := e.DeleteLocationEndpoint(ctx, req)
	if err != nil {
		return payload.DeleteLocationResponse{err.Error(), 500}, nil
	}
	getResp := resp.(payload.DeleteLocationResponse)
	return getResp, nil
}
