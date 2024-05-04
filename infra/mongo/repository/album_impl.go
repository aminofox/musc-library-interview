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

type albumRepositoryImpl struct {
	mongodb *mongodb2.MongoDB
}

const CollectionAlbum = "albums"

func NewAlbumRepository(mongodb *mongodb2.MongoDB) repository.AlbumRepository {
	return &albumRepositoryImpl{mongodb}
}

func (u *albumRepositoryImpl) Create(ctx context.Context, album *entity.Album) (string, error) {
	collection := u.mongodb.Collection(CollectionAlbum)
	res, err := collection.InsertOne(ctx, album)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *albumRepositoryImpl) Update(ctx context.Context, albumID string, updatedAlbum *entity.Album) (int64, error) {
	collection := u.mongodb.Collection(CollectionAlbum)
	objID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updatedAlbum,
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (u *albumRepositoryImpl) GetByID(ctx context.Context, albumID string) (*entity.Album, error) {
	var album *entity.Album
	collection := u.mongodb.Collection(CollectionAlbum)
	objID, err := primitive.ObjectIDFromHex(albumID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&album)
	return album, err
}

func (u *albumRepositoryImpl) GetList(ctx context.Context, params entity.GetListAlbumOption) ([]*entity.Album, error) {
	var albums []*entity.Album
	collection := u.mongodb.Collection(CollectionAlbum)

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
			bson.M{"title": params.Title},
			bson.M{"release_year": params.ReleaseYear},
			bson.M{"tracks": params.Track},
		},
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var album *entity.Album
		if err := cursor.Decode(&album); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (u *albumRepositoryImpl) Delete(ctx context.Context, albumID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionAlbum)
	objID, err := primitive.ObjectIDFromHex(albumID)
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
