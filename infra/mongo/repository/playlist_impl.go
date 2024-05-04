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

type playlistRepository struct {
	mongodb *mongodb2.MongoDB
}

const CollectionPlaylist = "playlists"

func NewPlaylistRepository(mongodb *mongodb2.MongoDB) repository.PlaylistRepository {
	return &playlistRepository{mongodb}
}

func (u *playlistRepository) Create(ctx context.Context, playlist *entity.Playlist) (string, error) {
	collection := u.mongodb.Collection(CollectionPlaylist)
	res, err := collection.InsertOne(ctx, playlist)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *playlistRepository) Update(ctx context.Context, playlistID string, updatedPlaylist *entity.Playlist) (int64, error) {
	collection := u.mongodb.Collection(CollectionPlaylist)
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return 0, err
	}
	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": updatedPlaylist,
	}
	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (u *playlistRepository) GetByID(ctx context.Context, playlistID string) (*entity.Playlist, error) {
	var playlist *entity.Playlist
	collection := u.mongodb.Collection(CollectionPlaylist)
	objID, err := primitive.ObjectIDFromHex(playlistID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&playlist)
	return playlist, err
}

func (u *playlistRepository) GetList(ctx context.Context, params entity.GetListPlaylistOption) ([]*entity.Playlist, error) {
	var playlists []*entity.Playlist
	collection := u.mongodb.Collection(CollectionPlaylist)

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
			bson.M{"tracks": params.Track},
		},
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var playlist *entity.Playlist
		if err := cursor.Decode(&playlist); err != nil {
			return nil, err
		}
		playlists = append(playlists, playlist)
	}
	return playlists, nil
}

func (u *playlistRepository) Delete(ctx context.Context, playlistID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionPlaylist)
	objID, err := primitive.ObjectIDFromHex(playlistID)
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
