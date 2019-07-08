package service

import (
	"context"
)

//UsersService for get user
type UsersService interface {
	GetUsers(ctx context.Context) (string, error)
}

type usersService struct{}

//NewUsersService for
func NewUsersService() UsersService {
	return usersService{}
}
func (usersService) GetUsers(ctx context.Context) (string, error) {
	return "Users Succesfully Loaded", nil
}
