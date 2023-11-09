package gapi

import (
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/pb"
	"github.com/mas-wig/post-api-1/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	authServices   services.AuthService
	userServices   services.UserService
	userCollection *mongo.Collection
	config         config.Config
}

func NewGRPCSever(
	config config.Config,
	authServices services.AuthService,
	userServices services.UserService,
	userCollection *mongo.Collection,
) (*Server, error) {
	return &Server{
		config:         config,
		authServices:   authServices,
		userServices:   userServices,
		userCollection: userCollection,
	}, nil
}
