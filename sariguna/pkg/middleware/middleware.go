package middleware

import (
	"sariguna_backend/sariguna/pkg/helpers"
	"sariguna_backend/sariguna/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")

	if tokenHeader == "" {
		c.JSON(401, helpers.ErrorResponse("request does not contain an access token"))
		c.Abort()
		return
	}

	tokenHeader = jwt.ExtractToken(tokenHeader)

	userId, err := jwt.VerifyToken(tokenHeader)

	if err != nil {
		c.JSON(401, helpers.ErrorResponse(err.Error()))
		c.Abort()
		return
	}

	c.Set("userId", userId)
}
