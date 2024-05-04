package auth

import (
	"strings"

	"music-libray-management/domain/entity"
	jwtPkg "music-libray-management/internal/jwt"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"e" binding:"required,email,IsValidEmail"`
	Password string `json:"p" binding:"required,max=16"`
}

type LoginUser struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
}

type LoginOutput struct {
	User         LoginUser `json:"user"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

func (au *authUseCase) Login(ctx *gin.Context, input *LoginInput) (*LoginOutput, error) {
	var err error
	var user *entity.User

	user, err = au.userRepository.GetUserByEmail(ctx, strings.ToLower(input.Email))
	if err != nil {
		return nil, ErrUserNotFound
	}

	err = au.passwordService.CheckHashPassword(*user.Password, input.Password)
	if err != nil {
		return nil, ErrUserNotFound
	}

	accessToken, err := au.jwtService.GenerateAccessToken(&jwtPkg.GenerateTokenInput{
		ID:    user.ID.Hex(),
		Email: user.Email,
	})
	if err != nil {
		return nil, ErrGenAccessTokenFailed
	}
	refreshToken, err := au.jwtService.GenerateRefreshToken(&jwtPkg.GenerateTokenInput{
		ID:    user.ID.Hex(),
		Email: user.Email,
	})

	if err != nil {
		return nil, ErrGenRefreshTokenFailed
	}

	return &LoginOutput{
		LoginUser{
			ID:       user.ID.Hex(),
			FullName: user.FullName,
		},
		accessToken,
		refreshToken,
	}, nil
}
