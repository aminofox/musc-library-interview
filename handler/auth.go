package handler

import (
	"music-libray-management/internal/response"
	"music-libray-management/usecase/auth"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, authUseCase auth.AuthUseCase, responseService response.ServiceResponse) {
	var input auth.LoginInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := authUseCase.Login(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func Register(ctx *gin.Context, authUseCase auth.AuthUseCase, responseService response.ServiceResponse) {
	var input auth.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := authUseCase.Register(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}

func RefreshToken(ctx *gin.Context, authUseCase auth.AuthUseCase, responseService response.ServiceResponse) {
	var input auth.RefreshTokenInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		responseService.AbortWithAppBindingError(ctx, err)
		return
	}
	output, err := authUseCase.RefreshToken(ctx, &input)
	if err != nil {
		responseService.AbortWithAppSingleError(ctx, err)
		return
	}
	responseService.ResponseSuccess(ctx, output, true)
}
