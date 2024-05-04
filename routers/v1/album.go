package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/album"

	"github.com/gin-gonic/gin"
)

func initAlbumRouter(
	r gin.IRouter,
	artistUseCase album.UseCase,
	responseService response.ServiceResponse,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreateAlbum(ctx, artistUseCase, responseService)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListAlbum(ctx, artistUseCase, responseService)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetAlbum(ctx, artistUseCase, responseService)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdateAlbum(ctx, artistUseCase, responseService)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeleteAlbum(ctx, artistUseCase, responseService)
	})
}
