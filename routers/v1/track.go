package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/track"

	"github.com/gin-gonic/gin"
)

func initTrackRouter(
	r gin.IRouter,
	trackUseCase track.UseCase,
	responseService response.ServiceResponse,
) {
	r.POST("", func(ctx *gin.Context) {
		handler.CreateTrack(ctx, trackUseCase, responseService)
	})
	r.GET("", func(ctx *gin.Context) {
		handler.GetListTrack(ctx, trackUseCase, responseService)
	})
	r.GET("/:id", func(ctx *gin.Context) {
		handler.GetTrack(ctx, trackUseCase, responseService)
	})
	r.PUT("/:id", func(ctx *gin.Context) {
		handler.UpdateTrack(ctx, trackUseCase, responseService)
	})
	r.DELETE("/:id", func(ctx *gin.Context) {
		handler.DeleteTrack(ctx, trackUseCase, responseService)
	})
}
