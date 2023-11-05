package services

import "github.com/mas-wig/post-api-1/types"

type UserService interface {
	FindUserByID(id string) (*types.DBResponse, error)
	FindUserByEmail(email string) (*types.DBResponse, error)
	UpdateUserByID(id string, data *types.UpdateInput) (*types.DBResponse, error)
}
