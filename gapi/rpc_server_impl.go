package gapi

import (
	"context"
	"strings"
	"time"

	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/pb"
	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpUser(ctx context.Context, req *pb.SignUpUserInput) (
	*pb.GenericResponse, error,
) {
	if req.GetPassword() != req.GetPasswordConfirm() {
		return nil, status.Errorf(codes.InvalidArgument, "password do not match")
	}
	user := types.RegisterInput{
		Name:            req.GetEmail(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
	}
	newUser, err := s.authServices.RegisterUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "email already exist") {
			return nil, status.Errorf(codes.AlreadyExists, "%s", err.Error())
		}
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	var (
		randomCode       = randstr.String(20)
		VerificationCode = utils.Encode(randomCode)
		updateData       = &types.UpdateInput{VerificationCode: VerificationCode}
		firstName        = newUser.Name
	)

	s.userServices.UpdateUserByID(newUser.ID.Hex(), updateData)
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	config, _ := config.LoadConfig("..")
	emailData := utils.EmailData{
		URL:       config.Origin + "/api/auth/verify-email/" + randomCode,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	if err = utils.SendEmail(newUser, &emailData, "verification_code.html"); err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}
	return &pb.GenericResponse{
		Status:  "success",
		Message: "we send email and verification code to " + newUser.Email,
	}, nil
}

func (s *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.GenericResponse, error) {
	var (
		randomCode       = req.GetVerificationCode()
		verificationCode = utils.Encode(randomCode)
		query            = bson.D{{Key: "verificationCode", Value: verificationCode}}
		updateDate       = bson.D{
			{Key: "$set", Value: bson.D{{Key: "verified", Value: true}, {Key: "updated_at", Value: time.Now()}}},
			{Key: "$unset", Value: bson.D{{Key: "verificationCode", Value: ""}}},
		}
	)

	result, err := s.userCollection.UpdateOne(ctx, query, updateDate)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	if result.MatchedCount == 0 {
		return nil, status.Errorf(codes.PermissionDenied, "could not verify email address")
	}

	return &pb.GenericResponse{Status: "success", Message: "email verified successfully"}, nil
}

func (s *Server) SignInUser(ctx context.Context, req *pb.SignInUserInput) (*pb.SignInUserResponse, error) {
	user, err := s.userServices.FindUserByEmail(req.GetEmail())
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Errorf(codes.InvalidArgument, "invalid email or password")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if !user.Verified {
		return nil, status.Errorf(codes.PermissionDenied, "your are not verified, please verify your email to login")
	}
	return &pb.SignInUserResponse{}, nil
}
