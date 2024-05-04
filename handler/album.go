package handler

import (
	"music-libray-management/internal/pagination"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/album"

	"github.com/gin-gonic/gin"
)

func CreateAlbum(ctx *gin.Context, albumUseCase album.UseCase, responseService response.ServiceResponse) {
	var input album.CreateAlbumInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := albumUseCase.Create(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetAlbum(ctx *gin.Context, albumUseCase album.UseCase, responseService response.ServiceResponse) {
	var input album.GetAlbumByIDInput
	input.ID = ctx.Param("id")
	output, err := albumUseCase.GetByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetListAlbum(ctx *gin.Context, albumUseCase album.UseCase, responseService response.ServiceResponse) {
	var input album.GetListAlbumInput
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
	output, err := albumUseCase.GetListAlbum(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func UpdateAlbum(ctx *gin.Context, albumUseCase album.UseCase, responseService response.ServiceResponse) {
	var input album.UpdateAlbumInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	input.ID = ctx.Param("id")
	output, err := albumUseCase.Update(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func DeleteAlbum(ctx *gin.Context, albumUseCase album.UseCase, responseService response.ServiceResponse) {
	var input album.DeleteAlbumInput
	input.ID = ctx.Param("id")
	output, err := albumUseCase.DeleteByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
