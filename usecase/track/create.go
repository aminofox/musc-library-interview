package track

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CreateTrackInput struct {
	Title       string `json:"title" bson:"title" binding:"required"`
	Artist      string `json:"artist" bson:"artist" binding:"required"`
	Album       string `json:"album" bson:"album"`
	Genre       string `json:"genre" bson:"genre"`
	ReleaseYear int    `json:"release_year" bson:"release_year"`
	Duration    int    `json:"duration" bson:"duration" binding:"required"`
	Mp3File     string `json:"mp3_file" bson:"mp3_file" binding:"required"`
}

type CreateTrackOutput struct {
	ID string `json:"id"`
}

func (t *trackUseCase) Create(ctx *gin.Context, input *CreateTrackInput) (*CreateTrackOutput, error) {
	data := &entity.Track{}

	err := copier.Copy(data, input)
	if err != nil {
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	session, err := t.database.Client().StartSession()
	if err != nil {
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}
	defer session.EndSession(ctx)

	err = session.StartTransaction()
	if err != nil {
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	id, err := t.trackRepository.Create(ctx, data)
	if err != nil {
		session.AbortTransaction(ctx)
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	currentArtist, err := t.artistRepository.GetByID(ctx, input.Artist)
	if err != nil {
		session.AbortTransaction(ctx)
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	currentArtist.Tracks = append(currentArtist.Tracks, id)
	_, err = t.artistRepository.Update(ctx, input.Artist, currentArtist)
	if err != nil {
		session.AbortTransaction(ctx)
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	currentAlbum, err := t.albumRepository.GetByID(ctx, input.Album)
	if err != nil {
		session.AbortTransaction(ctx)
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	currentAlbum.Tracks = append(currentAlbum.Tracks, id)
	_, err = t.albumRepository.Update(ctx, input.Album, currentAlbum)
	if err != nil {
		session.AbortTransaction(ctx)
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return &CreateTrackOutput{ID: ""}, ErrCreateTrackFailed
	}

	return &CreateTrackOutput{
		ID: id,
	}, nil
}
