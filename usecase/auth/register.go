package auth

import (
	"strings"

	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type RegisterOutput struct {
	ID string `json:"id"`
}

func (au authUseCase) Register(ctx *gin.Context, input *RegisterInput) (*RegisterOutput, error) {
	if input.Password != input.ConfirmPassword {
		return nil, ErrPasswordConfirmPasswordDifference
	}

	existedUser, _ := au.userRepository.GetUserByEmail(ctx, input.Email)
	if existedUser != nil {
		return nil, ErrUserExisted
	}

	hashPassword, err := au.passwordService.HashPassword(input.Password)
	if err != nil {
		return nil, ErrGenPasswordFailed
	}

	parts := strings.Split(input.Email, "@")
	fullName := parts[0]

	id, err := au.userRepository.Create(ctx, &entity.User{
		FullName: fullName,
		Email:    input.Email,
		Password: &hashPassword,
	})

	if err != nil {
		return nil, ErrRegisterFailed
	}

	return &RegisterOutput{ID: id}, nil
}
