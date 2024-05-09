package repository

import (
	"context"
	"music-libray-management/domain/entity"
	"music-libray-management/domain/repository"
	mongodb2 "music-libray-management/infra/mongo"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type artistRepositoryImpl struct {
	mongodb *mongodb2.MongoDB
}

const CollectionArtist = "artists"

func NewArtistRepository(mongodb *mongodb2.MongoDB) repository.ArtistRepository {
	return &artistRepositoryImpl{mongodb}
}

func (u *artistRepositoryImpl) Create(ctx context.Context, artist *entity.Artist) (string, error) {
	collection := u.mongodb.Collection(CollectionArtist)
	res, err := collection.InsertOne(ctx, artist)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *artistRepositoryImpl) Update(ctx context.Context, artistID string, updatedArtist *entity.Artist) (int64, error) {
	collection := u.mongodb.Collection(CollectionArtist)
	objID, err := primitive.ObjectIDFromHex(artistID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updatedArtist,
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (u *artistRepositoryImpl) GetByID(ctx context.Context, artistID string) (*entity.Artist, error) {
	var artist *entity.Artist
	collection := u.mongodb.Collection(CollectionArtist)
	objID, err := primitive.ObjectIDFromHex(artistID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&artist)
	return artist, err
}

func (u *artistRepositoryImpl) GetByName(ctx context.Context, name string) ([]string, error) {
	var artists []string
	collection := u.mongodb.Collection(CollectionArtist)
	filter := bson.M{"name": bson.M{"$regex": name, "$options": "i"}}
	cursor, err := collection.Find(ctx, filter)
	for cursor.Next(ctx) {
		var artist *entity.Artist
		if err := cursor.Decode(&artist); err != nil {
			return nil, err
		}
		artists = append(artists, artist.ID.Hex())
	}

	return artists, err
}

func (u *artistRepositoryImpl) GetList(ctx context.Context, params entity.GetListArtistOption) ([]*entity.Artist, error) {
	var artists []*entity.Artist
	collection := u.mongodb.Collection(CollectionArtist)

	findOptions := &options.FindOptions{}
	if params.Order != "" {
		arr := strings.Split(params.Order, " ")
		findOptions.SetSort(bson.D{{arr[0], arr[1]}})
	}

	if params.PageIndex > 0 {
		findOptions.SetLimit(int64(params.PageSize))
		findOptions.SetSkip(int64(params.PageSize * (params.PageIndex - 1)))
	}

	filter := bson.M{
		"$or": []interface{}{
			bson.M{"name": params.Name},
			bson.M{"country": params.Country},
			bson.M{"tracks": params.Track},
		},
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var artist *entity.Artist
		if err := cursor.Decode(&artist); err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}

func (u *artistRepositoryImpl) Delete(ctx context.Context, artistID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionArtist)
	objID, err := primitive.ObjectIDFromHex(artistID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objID}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}
