package repository

import (
	"context"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	mongodb2 "music-libray-management/infra/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepositoryImpl struct {
	mongodb *mongodb2.MongoDB
}

const CollectionUser = "users"

func NewUserRepository(mongodb *mongodb2.MongoDB) repository.UserRepository {
	return &userRepositoryImpl{mongodb}
}

func (u *userRepositoryImpl) Create(ctx context.Context, user *entity.User) (string, error) {
	collection := u.mongodb.Collection(CollectionUser)
	res, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *userRepositoryImpl) Update(ctx context.Context, userID string, updatedUser *entity.User) (string, error) {
	collection := u.mongodb.Collection(CollectionUser)
	filter := bson.M{"_id": userID}
	update := bson.M{
		"$set": bson.M{
			"full_name": updatedUser.FullName,
			"password":  updatedUser.Password,
		},
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", err
	}

	return res.UpsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *userRepositoryImpl) GetByID(ctx context.Context, userID string) (*entity.User, error) {
	var user *entity.User
	collection := u.mongodb.Collection(CollectionUser)
	objID, err := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func (u *userRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user *entity.User
	collection := u.mongodb.Collection(CollectionUser)
	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func (u *userRepositoryImpl) GetList(ctx context.Context) ([]*entity.User, error) {
	var users []*entity.User
	collection := u.mongodb.Collection(CollectionUser)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user *entity.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *userRepositoryImpl) Delete(ctx context.Context, userID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionUser)
	objID, err := primitive.ObjectIDFromHex(userID)
	filter := bson.M{"_id": objID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
