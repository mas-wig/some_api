package services

import (
	"context"
	"strings"

	"github.com/mas-wig/post-api-1/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserServiceImpl(collection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{collection: collection, ctx: ctx}
}

func (us *UserServiceImpl) FindUserByID(id string) (*types.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var (
		user  *types.DBResponse
		query = bson.M{"_id": oid}
	)

	if err := us.collection.FindOne(us.ctx, &query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return &types.DBResponse{}, err
		}
		return nil, err
	}
	return user, nil
}

func (us *UserServiceImpl) FindUserByEmail(email string) (*types.DBResponse, error) {
	var (
		user  *types.DBResponse
		query = bson.M{"email": strings.ToLower(email)}
	)

	if err := us.collection.FindOne(us.ctx, query).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return &types.DBResponse{}, nil
		}
		return nil, err
	}
	return user, nil
}
