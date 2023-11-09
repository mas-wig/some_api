package gapi

import (
	"context"

	"github.com/mas-wig/post-api-1/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetMeRequest) (*pb.UserResponse, error) {
	id := req.GetId()
	user, err := s.userServices.FindUserByID(id)
	if err != nil {
		return nil, status.Errorf(codes.Unimplemented, err.Error())
	}
	return &pb.UserResponse{
		User: &pb.User{
			Id:        user.ID.Hex(),
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
