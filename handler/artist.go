package handler

import (
	"music-libray-management/internal/pagination"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/artist"

	"github.com/gin-gonic/gin"
)

func CreateArtist(ctx *gin.Context, artistUseCase artist.UseCase, responseService response.ServiceResponse) {
	var input artist.CreateArtistInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := artistUseCase.Create(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetArtist(ctx *gin.Context, artistUseCase artist.UseCase, responseService response.ServiceResponse) {
	var input artist.GetArtistByIDInput
	input.ID = ctx.Param("id")
	output, err := artistUseCase.GetByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetListArtist(ctx *gin.Context, artistUseCase artist.UseCase, responseService response.ServiceResponse) {
	var input artist.GetListArtistInput
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
	output, err := artistUseCase.GetListArtist(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func UpdateArtist(ctx *gin.Context, artistUseCase artist.UseCase, responseService response.ServiceResponse) {
	var input artist.UpdateArtistInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	input.ID = ctx.Param("id")
	output, err := artistUseCase.Update(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func DeleteArtist(ctx *gin.Context, artistUseCase artist.UseCase, responseService response.ServiceResponse) {
	var input artist.DeleteArtistInput
	input.ID = ctx.Param("id")
	output, err := artistUseCase.DeleteByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
