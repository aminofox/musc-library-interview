package handler

import (
	"music-libray-management/internal/response"
	"music-libray-management/middlewares"
	"music-libray-management/usecase/document"

	"github.com/gin-gonic/gin"
)

func UploadDocument(ctx *gin.Context, documentUseCase document.UseCase, responseService response.ServiceResponse) {
	var input document.UploadFileInput
	if err := ctx.ShouldBind(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	user := middlewares.MustGetUser(ctx)
	input.UserID = user.ID.Hex()
	output, err := documentUseCase.UploadFile(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
