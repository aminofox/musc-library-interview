package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/middlewares"
	"music-libray-management/usecase/auth"

	"github.com/gin-gonic/gin"
)

func initAuthRouter(
	r gin.IRouter,
	authUseCase auth.AuthUseCase,
	middleware middlewares.Middleware,
	responseService response.ServiceResponse,
) {
	r.POST("/login", func(context *gin.Context) {
		handler.Login(context, authUseCase, responseService)
	})

	r.POST("/register", func(context *gin.Context) {
		handler.Register(context, authUseCase, responseService)
	})

	r.POST("/refresh-token", func(context *gin.Context) {
		handler.RefreshToken(context, authUseCase, responseService)
	})
}
