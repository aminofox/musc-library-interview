package track

import (
	"music-libray-management/domain/entity"
	"music-libray-management/internal/helper"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UpdateTrackInput struct {
	ID          string
	Title       string `json:"title" bson:"title"`
	Artist      string `json:"artist" bson:"artist"`
	Album       string `json:"album" bson:"album"`
	Genre       string `json:"genre" bson:"genre"`
	ReleaseYear int    `json:"release_year" bson:"release_year"`
	Duration    int    `json:"duration" bson:"duration"`
	Mp3File     string `json:"mp3_file" bson:"mp3_file"`
}

type UpdateTrackOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *trackUseCase) Update(ctx *gin.Context, input *UpdateTrackInput) (*UpdateTrackOutput, error) {
	data := &entity.Track{}

	err := copier.Copy(data, input)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	currentTrack, err := t.trackRepository.GetByID(ctx, input.ID)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	currentAlbum, err := t.albumRepository.GetByID(ctx, currentTrack.Album)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	currentArtist, err := t.artistRepository.GetByID(ctx, currentTrack.Artist)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	newAlbum, err := t.albumRepository.GetByID(ctx, input.Album)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	newArtist, err := t.artistRepository.GetByID(ctx, input.Artist)
	if err != nil {
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	session, err := t.database.Client().StartSession()
	if err != nil {
		return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
	}
	defer session.EndSession(ctx)

	err = session.StartTransaction()
	if err != nil {
		return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
	}

	if currentTrack.Album != input.Album {
		newAlbum.Tracks = append(newAlbum.Tracks, input.ID)
		currentAlbum.Tracks = helper.RemoveElement(input.ID, currentAlbum.Tracks)

		_, err := t.albumRepository.Update(ctx, newAlbum.ID.Hex(), newAlbum)
		if err != nil {
			session.AbortTransaction(ctx)
			return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
		}

		_, err = t.albumRepository.Update(ctx, currentAlbum.ID.Hex(), currentAlbum)
		if err != nil {
			session.AbortTransaction(ctx)
			return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
		}
	}

	if currentTrack.Artist != input.Artist {
		newArtist.Tracks = append(newArtist.Tracks, input.ID)
		currentArtist.Tracks = helper.RemoveElement(input.ID, currentArtist.Tracks)

		_, err := t.artistRepository.Update(ctx, newArtist.ID.Hex(), newArtist)
		if err != nil {
			session.AbortTransaction(ctx)
			return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
		}

		_, err = t.artistRepository.Update(ctx, currentArtist.ID.Hex(), currentArtist)
		if err != nil {
			session.AbortTransaction(ctx)
			return &UpdateTrackOutput{0}, ErrUpdateTrackFailed
		}
	}

	rowsAffected, err := t.trackRepository.Update(ctx, input.ID, data)
	if err != nil {
		session.AbortTransaction(ctx)
		return &UpdateTrackOutput{RowsAffected: 0}, ErrUpdateTrackFailed
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return &UpdateTrackOutput{0}, ErrCreateTrackFailed
	}

	return &UpdateTrackOutput{
		rowsAffected,
	}, nil
}
