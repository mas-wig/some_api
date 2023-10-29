package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterInput struct {
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
	Name            string    `json:"name" bson:"name" binding:"required"`
	Email           string    `json:"email" bson:"email" binding:"required"`
	Password        string    `json:"password" bson:"password" binding:"required"`
	PasswordConfirm string    `json:"passwordConfirm" bson:"passwordConfirm" binding:"required"`
	Role            string    `json:"role" bson:"role" binding:"required"`
	Verified        bool      `json:"verified" bson:"verified" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type DBResponse struct {
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name            string             `json:"name" bson:"name" binding:"required"`
	Email           string             `json:"email" bson:"email" binding:"required"`
	Password        string             `json:"password" bson:"password" binding:"required"`
	PasswordConfirm string             `json:"passwordConfirm" bson:"passwordConfirm" binding:"required"`
	Role            string             `json:"role" bson:"role" binding:"required"`
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	Verified        bool               `json:"verified" bson:"verified" binding:"required"`
}

type UserResponse struct {
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	Name      string             `json:"name,omitempty" bson:"name" binding:"required"`
	Email     string             `json:"email,omitempty" bson:"email" binding:"required"`
	Role      string             `json:"role,omitempty" bson:"role" binding:"required"`
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Verified  bool               `json:"verified" bson:"verified" binding:"required"`
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
