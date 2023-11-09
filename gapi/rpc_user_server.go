package gapi

import (
	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/pb"
	"github.com/mas-wig/post-api-1/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	config         config.Config
	userServices   services.UserService
	userCollection *mongo.Collection
}

func NewGrpcUserServer(
	config config.Config, userService services.UserService, userCollection *mongo.Collection,
) (*UserServer, error) {
	return &UserServer{config: config, userServices: userService, userCollection: userCollection}, nil
}
