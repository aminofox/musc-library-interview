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

type trackRepositoryImpl struct {
	mongodb *mongodb2.MongoDB
}

const CollectionTrack = "tracks"

func NewTrackRepository(mongodb *mongodb2.MongoDB) repository.TrackRepository {
	return &trackRepositoryImpl{mongodb}
}

func (u *trackRepositoryImpl) Create(ctx context.Context, track *entity.Track) (string, error) {
	collection := u.mongodb.Collection(CollectionTrack)
	res, err := collection.InsertOne(ctx, track)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (u *trackRepositoryImpl) Update(ctx context.Context, trackID string, updatedTrack *entity.Track) (int64, error) {
	collection := u.mongodb.Collection(CollectionTrack)
	objID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{}}

	if updatedTrack.Title != "" {
		update["$set"].(bson.M)["title"] = updatedTrack.Title
	}

	if updatedTrack.Artist != "" {
		update["$set"].(bson.M)["artist"] = updatedTrack.Artist
	}

	if updatedTrack.Album != "" {
		update["$set"].(bson.M)["album"] = updatedTrack.Album
	}

	if updatedTrack.Genre != "" {
		update["$set"].(bson.M)["genre"] = updatedTrack.Genre
	}

	if updatedTrack.Duration != 0 {
		update["$set"].(bson.M)["duration"] = updatedTrack.Duration
	}

	if updatedTrack.ReleaseYear != 0 {
		update["$set"].(bson.M)["release_year"] = updatedTrack.ReleaseYear
	}

	if updatedTrack.Mp3File != "" {
		update["$set"].(bson.M)["mp3_file"] = updatedTrack.Mp3File
	}

	res, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	return res.MatchedCount, nil
}

func (u *trackRepositoryImpl) GetByID(ctx context.Context, trackID string) (*entity.Track, error) {
	var track *entity.Track
	collection := u.mongodb.Collection(CollectionTrack)
	objID, err := primitive.ObjectIDFromHex(trackID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	err = collection.FindOne(ctx, filter).Decode(&track)
	return track, err
}

func (u *trackRepositoryImpl) GetList(ctx context.Context, params entity.GetListTrackOption) ([]*entity.Track, error) {
	var tracks []*entity.Track
	collection := u.mongodb.Collection(CollectionTrack)

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
			bson.M{"artist": params.Artist},
			bson.M{"album": params.Album},
			bson.M{"genre": params.Genre},
		},
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var track *entity.Track
		if err := cursor.Decode(&track); err != nil {
			return nil, err
		}
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (u *trackRepositoryImpl) Delete(ctx context.Context, trackID string) (int64, error) {
	collection := u.mongodb.Collection(CollectionTrack)
	objID, err := primitive.ObjectIDFromHex(trackID)
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
