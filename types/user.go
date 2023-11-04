package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterInput struct {
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" bson:"updated_at"`
	Name            string    `json:"name" bson:"name" binding:"required"`
	Email           string    `json:"email" bson:"email" binding:"required"`
	Password        string    `json:"password" bson:"password" binding:"required,min=8"`
	PasswordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm,omitempty" binding:"required"`
	Role            string    `json:"role" bson:"role"`
	Verified        bool      `json:"verified" bson:"verified"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DBResponse struct {
	CreatedAt       time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at"`
	Name            string             `json:"name" bson:"name"`
	Email           string             `json:"email" bson:"email"`
	Password        string             `json:"password" bson:"password"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm"`
	Role            string             `json:"role" bson:"role"`
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Verified        bool               `json:"verified" bson:"verified"`
}

type UserResponse struct {
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Name      string             `json:"name,omitempty" bson:"name"`
	Email     string             `json:"email,omitempty" bson:"email"`
	Role      string             `json:"role,omitempty" bson:"role"`
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Verified  bool               `json:"verified" bson:"verified"`
}

// Jadi gak semata mata langsung ontput semua data dari database
func FilteredResponse(user *DBResponse) UserResponse {
	return UserResponse{
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		ID:        user.ID,
	}
}
