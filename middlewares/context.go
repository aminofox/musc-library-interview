package middlewares

import (
	"music-libray-management/domain/entity"

	"github.com/gin-gonic/gin"
)

func MustGetUser(c *gin.Context) *entity.User {
	user, _ := c.Get("AuthorizedUser")
	out := user.(*entity.User)
	if user == nil {
		panic("unauthorized")
	}
	return out
}
