package v1

import (
	"music-libray-management/handler"
	"music-libray-management/internal/response"
	"music-libray-management/usecase/document"

	"github.com/gin-gonic/gin"
)

func initDocumentRouter(
	r gin.IRouter,
	useCase document.UseCase,
	responseService response.ServiceResponse,
) {

	r.POST("/file-upload", func(ctx *gin.Context) {
		handler.UploadDocument(ctx, useCase, responseService)
	})

}
