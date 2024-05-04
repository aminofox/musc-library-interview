package v1

import (
	"music-libray-management/internal/response"
	"music-libray-management/middlewares"
	"music-libray-management/usecase/album"
	"music-libray-management/usecase/artist"
	"music-libray-management/usecase/auth"
	"music-libray-management/usecase/document"
	"music-libray-management/usecase/playlist"
	"music-libray-management/usecase/track"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitV1Router(
	r *gin.RouterGroup,
	middleware middlewares.Middleware,
	authUseCase auth.AuthUseCase,
	trackUseCase track.UseCase,
	documentUseCase document.UseCase,
	playlistUseCase playlist.UseCase,
	artistUseCase artist.UseCase,
	albumUseCase album.UseCase,
	responseService response.ServiceResponse,
) {
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong v1",
		})
	})

	r.Use()
	{
		initAuthRouter(r.Group("/auth"), authUseCase, middleware, responseService)
	}

	r.Use(middleware.Authentication())
	{
		initTrackRouter(r.Group("/track"), trackUseCase, responseService)
		initDocumentRouter(r.Group("/document"), documentUseCase, responseService)
		initPlaylistRouter(r.Group("/playlist"), playlistUseCase, responseService)
		initArtistRouter(r.Group("/artist"), artistUseCase, responseService)
		initAlbumRouter(r.Group("/album"), albumUseCase, responseService)
	}
}
