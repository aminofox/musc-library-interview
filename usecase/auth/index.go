package auth

import (
	"errors"
	"music-libray-management/config"
	"music-libray-management/domain/repository"
	"music-libray-management/infra/mongo"
	jwtPkg "music-libray-management/internal/jwt"
	passwordPkg "music-libray-management/internal/password"

	"github.com/gin-gonic/gin"
)

var (
	ErrLoginFailed                       = errors.New("login failed")
	ErrRegisterFailed                    = errors.New("register failed")
	ErrPasswordConfirmPasswordDifference = errors.New("password must be same confirm password")
	ErrUserExisted                       = errors.New("user existed")
	ErrGenPasswordFailed                 = errors.New("password invalid")
	ErrGenAccessTokenFailed              = errors.New(("token failed"))
	ErrGenRefreshTokenFailed             = errors.New(("refresh token failed"))
	ErrUserNotFound                      = errors.New("user not existed")
)

type AuthUseCase interface {
	Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, error)
	Register(ctx *gin.Context, input *RegisterInput) (*RegisterOutput, error)
	RefreshToken(ctx *gin.Context, input *RefreshTokenInput) (*RefreshTokenOutput, error)
}

type authUseCase struct {
	cfg             *config.Environment
	jwtService      jwtPkg.Service
	passwordService passwordPkg.Service
	userRepository  repository.UserRepository
	database        *mongo.MongoDB
}

func NewAuthUseCase(
	cfg *config.Environment,
	jwtService jwtPkg.Service,
	passwordService passwordPkg.Service,
	userRepository repository.UserRepository,
	database *mongo.MongoDB,
) AuthUseCase {
	return &authUseCase{
		cfg:             cfg,
		jwtService:      jwtService,
		passwordService: passwordService,
		userRepository:  userRepository,
		database:        database,
	}
}
