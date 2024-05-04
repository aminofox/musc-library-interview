package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/artist"

	"github.com/gin-gonic/gin"
)

func initArtistRouter(
	r gin.IRouter,
	artistUseCase artist.UseCase,
	responseService response.ServiceResponse,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreateArtist(ctx, artistUseCase, responseService)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListArtist(ctx, artistUseCase, responseService)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetArtist(ctx, artistUseCase, responseService)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdateArtist(ctx, artistUseCase, responseService)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeleteArtist(ctx, artistUseCase, responseService)
	})
}
