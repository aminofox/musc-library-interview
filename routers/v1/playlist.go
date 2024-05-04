package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/playlist"

	"github.com/gin-gonic/gin"
)

func initPlaylistRouter(
	r gin.IRouter,
	playlistUseCase playlist.UseCase,
	responseService response.ServiceResponse,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreatePlaylist(ctx, playlistUseCase, responseService)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListPlaylist(ctx, playlistUseCase, responseService)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetPlaylist(ctx, playlistUseCase, responseService)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdatePlaylist(ctx, playlistUseCase, responseService)
	})
	r.PUT("/add-track", func(ctx *gin.Context) {
		handler.AddTrackToPlaylist(ctx, playlistUseCase, responseService)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeletePlaylist(ctx, playlistUseCase, responseService)
	})
}
