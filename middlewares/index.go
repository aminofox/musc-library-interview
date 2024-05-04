package middlewares

import (
	"music-libray-management/domain/repository"
	"music-libray-management/internal/jwt"

	"github.com/gin-gonic/gin"
)

type Middleware interface {
	Authentication() gin.HandlerFunc
}

type middleware struct {
	jwtService     jwt.Service
	userRepository repository.UserRepository
}

func NewMiddleware(jwtService jwt.Service, userRepository repository.UserRepository) Middleware {
	return &middleware{
		jwtService:     jwtService,
		userRepository: userRepository,
	}
}
