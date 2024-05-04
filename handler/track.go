package handler

import (
	"music-libray-management/internal/pagination"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/track"

	"github.com/gin-gonic/gin"
)

func CreateTrack(ctx *gin.Context, trackUseCase track.UseCase, responseService response.ServiceResponse) {
	var input track.CreateTrackInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := trackUseCase.Create(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetTrack(ctx *gin.Context, trackUseCase track.UseCase, responseService response.ServiceResponse) {
	var input track.GetTrackByIDInput
	input.ID = ctx.Param("id")
	output, err := trackUseCase.GetByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetListTrack(ctx *gin.Context, trackUseCase track.UseCase, responseService response.ServiceResponse) {
	var input track.GetListTrackInput
	if err := ctx.ShouldBind(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	pageIndex, pageSize, order := pagination.GetDefaultPagination(
		input.PageIndex, input.PageSize, input.Order,
	)
	input.PageIndex = pageIndex
	input.PageSize = pageSize
	input.Order = order
	output, err := trackUseCase.GetListTrack(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func UpdateTrack(ctx *gin.Context, trackUseCase track.UseCase, responseService response.ServiceResponse) {
	var input track.UpdateTrackInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	input.ID = ctx.Param("id")
	output, err := trackUseCase.Update(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func DeleteTrack(ctx *gin.Context, trackUseCase track.UseCase, responseService response.ServiceResponse) {
	var input track.DeleteTrackInput
	input.ID = ctx.Param("id")
	output, err := trackUseCase.DeleteByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
