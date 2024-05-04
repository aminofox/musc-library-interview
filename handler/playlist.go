package handler

import (
	"music-libray-management/internal/pagination"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/playlist"

	"github.com/gin-gonic/gin"
)

func CreatePlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.CreatePlaylistInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := playlistUseCase.Create(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetPlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.GetPlaylistByIDInput
	input.ID = ctx.Param("id")
	output, err := playlistUseCase.GetByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func GetListPlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.GetListPlaylistInput
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
	output, err := playlistUseCase.GetListPlaylist(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func UpdatePlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.UpdatePlaylistInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	input.ID = ctx.Param("id")
	output, err := playlistUseCase.Update(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func DeletePlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.DeletePlaylistInput
	input.ID = ctx.Param("id")
	output, err := playlistUseCase.DeleteByID(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func AddTrackToPlaylist(ctx *gin.Context, playlistUseCase playlist.UseCase, responseService response.ServiceResponse) {
	var input playlist.AddTrackToPlaylistInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := playlistUseCase.AddTrackToPlaylist(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
