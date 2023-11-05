package services

import (
	"context"
	"errors"
	"strings"

	"github.com/mas-wig/post-api-1/types"
	"github.com/mas-wig/post-api-1/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (us *UserServiceImpl) UpdateUserByID(id string, data *types.UpdateInput) (*types.DBResponse, error) {
	doc, err := utils.ToDocument(data)
	if err != nil {
		return &types.DBResponse{}, err
	}

	objID, _ := primitive.ObjectIDFromHex(id)
	var (
		updatedUser *types.DBResponse
		query       = bson.D{{Key: "_id", Value: objID}}
		update      = bson.D{{Key: "$set", Value: doc}}
		result      = us.collection.FindOneAndUpdate(us.ctx, query, update, options.FindOneAndUpdate().SetReturnDocument(1))
	)
	if err := result.Decode(&updatedUser); err != nil {
		return nil, errors.New("no document with that id exists")
	}
	return updatedUser, nil
}
