package services

import (
	"errors"
	"strings"
	"time"

	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
)

type AuthServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewAuthServiceImpl(collection *mongo.Collection, ctx context.Context) *AuthServiceImpl {
	return &AuthServiceImpl{collection: collection, ctx: ctx}
}

func (as *AuthServiceImpl) LoginUser(user *types.LoginInput) (*types.DBResponse, error) {
	panic("not implemented")
}

func (as *AuthServiceImpl) RegisterUser(user *types.RegisterInput) (*types.DBResponse, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"

	hashedPass, _ := utils.HashedPassword(user.Password)
	user.Password = hashedPass

	insert, err := as.collection.InsertOne(as.ctx, &user)
	if err != nil {
		if err, ok := err.(mongo.WriteException); ok && err.WriteConcernError.Code == 11000 {
			return nil, errors.New("user with this email already exist")
		}
		return nil, err
	}

	var (
		newUser *types.DBResponse
		index   = mongo.IndexModel{Keys: bson.M{"email": 1}, Options: options.Index().SetUnique(true)}
		query   = bson.M{"_id": insert.InsertedID}
	)

	if _, err := as.collection.Indexes().CreateOne(as.ctx, index); err != nil {
		return nil, errors.New("could not create index for this email")
	}

	if err := as.collection.FindOne(as.ctx, query).Decode(&newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}
