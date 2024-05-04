package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m *middleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string

		headerAuthorization := ctx.GetHeader("Authorization")
		if headerAuthorization == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			ctx.Abort()
			return
		}

		if strings.HasPrefix(headerAuthorization, "Bearer ") {
			token = strings.TrimPrefix(headerAuthorization, "Bearer ")
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		verifiedToken, err := m.jwtService.ValidateAccessToken(token)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := m.userRepository.GetByID(ctx, verifiedToken.ID)
		if err != nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !user.IsValid() {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("AuthorizedUser", user)
		ctx.Next()
	}

}
