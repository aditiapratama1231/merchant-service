package endpoint

import (
	"context"
	"errors"

	model "qasir-supplier/marketplace/pkg/request/model" // part of transport
	"qasir-supplier/marketplace/pkg/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetUsersEndpoint(srv service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(model.GetUsersRequest) // we really just need the request, we don't use any value from it
		d, err := srv.GetUsers(ctx)
		if err != nil {
			return model.GetUsersResponse{d, err.Error()}, nil
		}
		return model.GetUsersResponse{d, ""}, nil
	}
}

func (e Endpoints) GetUsers(ctx context.Context) (string, error) {
	req := model.GetUsersRequest{}
	resp, err := e.GetUsersEndpoint(ctx, req)
	if err != nil {
		return "", err
	}
	getResp := resp.(model.GetUsersResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Users, nil
}
