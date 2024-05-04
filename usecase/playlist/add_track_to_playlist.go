package playlist

import (
	"music-libray-management/internal/helper"

	"github.com/gin-gonic/gin"
)

type AddTrackToPlaylistInput struct {
	PlaylistID string   `json:"playlist_id"`
	Tracks     []string `json:"tracks"`
}

type AddTrackToPlaylistOutput struct {
	RowsAffected int64 `json:"rowsAffected"`
}

func (t *playlistUseCase) AddTrackToPlaylist(ctx *gin.Context, input *AddTrackToPlaylistInput) (*AddTrackToPlaylistOutput, error) {
	currentPlaylist, err := t.playlistRepository.GetByID(ctx, input.PlaylistID)
	if err != nil {
		return nil, err
	}

	newTracks := append(currentPlaylist.Tracks, input.Tracks...)
	uniqueTracks := helper.RemoveDuplicate(newTracks)

	currentPlaylist.Tracks = uniqueTracks

	rowsAffected, err := t.playlistRepository.Update(ctx, input.PlaylistID, currentPlaylist)
	if err != nil {
		return &AddTrackToPlaylistOutput{RowsAffected: 0}, err
	}

	return &AddTrackToPlaylistOutput{
		RowsAffected: rowsAffected,
	}, nil
}
