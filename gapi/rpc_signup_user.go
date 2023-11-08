package gapi

import (
	"context"
	"strings"

	"github.com/mas-wig/post-api-1/config"
	"github.com/mas-wig/post-api-1/pb"
	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"github.com/thanhpk/randstr"
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
