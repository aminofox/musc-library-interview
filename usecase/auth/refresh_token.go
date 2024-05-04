package auth

import (
	jwtPkg "music-libray-management/internal/jwt"

	"github.com/gin-gonic/gin"
)

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RefreshTokenOutput struct {
	User         LoginUser `json:"user"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

func (au authUseCase) RefreshToken(ctx *gin.Context, input *RefreshTokenInput) (*RefreshTokenOutput, error) {
	verifyToken, err := au.jwtService.ValidateRefreshToken(input.RefreshToken)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user, err := au.userRepository.GetByID(ctx, verifyToken.ID)
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

	return &RefreshTokenOutput{
		LoginUser{
			ID:       user.ID.Hex(),
			FullName: user.FullName,
		},
		accessToken,
		refreshToken,
	}, nil
}
