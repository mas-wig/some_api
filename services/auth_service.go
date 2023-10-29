package services

import "github.com/mas-wig/post-api-1/types"

type AuthService interface {
	LoginUser(user *types.LoginInput) (*types.DBResponse, error)
	RegisterUser(user *types.RegisterInput) (*types.DBResponse, error)
}
