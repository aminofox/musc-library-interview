package routers

import (
	"music-libray-management/config"
	"music-libray-management/internal/response"
	"music-libray-management/middlewares"
	"music-libray-management/usecase/album"
	"music-libray-management/usecase/artist"
	"music-libray-management/usecase/auth"
	"music-libray-management/usecase/document"
	"music-libray-management/usecase/playlist"
	"music-libray-management/usecase/track"
	"music-libray-management/validations"
	"strings"
	"time"

	v1Routers "music-libray-management/routers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	API = "/api"
	V1  = "v1"
)

func InitRouter(
	cfg *config.Environment,
	middleware middlewares.Middleware,
	authUseCase auth.AuthUseCase,
	trackUseCase track.UseCase,
	documentUseCase document.UseCase,
	playlistUseCase playlist.UseCase,
	artistUseCase artist.UseCase,
	albumUseCase album.UseCase,
	responseService response.ServiceResponse,
) *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: strings.Split(cfg.AllowOrigins, ","),
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Headers",
			"Authorization",
			"X-XSRF-TOKEN",
		},
		ExposeHeaders: []string{
			"Content-Disposition",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("IsValidEmail", validations.IsValidEmail)
	}

	apiRouter := router.Group(API)

	v1Routers.InitV1Router(
		apiRouter.Group(V1),
		middleware,
		authUseCase,
		trackUseCase,
		documentUseCase,
		playlistUseCase,
		artistUseCase,
		albumUseCase,
		responseService,
	)

	return router
}
