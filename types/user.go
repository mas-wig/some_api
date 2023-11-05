package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterInput struct {
	ResetPasswordAt    time.Time `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
	Name               string    `json:"name" bson:"name" binding:"required"`
	Email              string    `json:"email" bson:"email" binding:"required"`
	Password           string    `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm    string    `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Role               string    `json:"role" bson:"role"`
	VerificationCode   string    `json:"verificationCode,omitempty" bson:"verificationCode,omitempty"`
	ResetPasswordToken string    `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	Verified           bool      `json:"verified" bson:"verified"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DBResponse struct {
	ResetPasswordAt    time.Time          `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
	Name               string             `json:"name" bson:"name"`
	Email              string             `json:"email" bson:"email"`
	Password           string             `json:"password" bson:"password"`
	PasswordConfirm    string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role               string             `json:"role" bson:"role"`
	VerificationCode   string             `json:"verificationCode,omitempty" bson:"verificationCode"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	Verified           bool               `json:"verified" bson:"verified"`
}

type UpdateInput struct {
	ResetPasswordAt    time.Time `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Name               string    `json:"name,omitempty" bson:"name,omitempty"`
	Email              string    `json:"email,omitempty" bson:"email,omitempty"`
	Password           string    `json:"password,omitempty" bson:"password,omitempty"`
	Role               string    `json:"role,omitempty" bson:"role,omitempty"`
	VerificationCode   string    `json:"verificationCode,omitempty" bson:"verificationCode,omitempty"`
	ResetPasswordToken string    `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	Verified           bool      `json:"verified,omitempty" bson:"verified,omitempty"`
}

type UserResponse struct {
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty"`
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
}

type ForgotPasswordInput struct {
	Email string `json:"email" bson:"email" binding:"required"`
}

type ResetPasswordInput struct {
	Password        string `json:"password" bson:"password"`
	PasswordConfirm string `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
}

func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
