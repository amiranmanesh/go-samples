package middleware

import (
	"awesome_webkits/database/models"
	"awesome_webkits/http/auth"
	"awesome_webkits/utils/parser"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userID uint
)

func RestAuthUserMiddleware(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")

	token, err := parser.ReadBearerToken(tokenHeader)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	userID, validErr := models.OauthAccessToken{}.VerifyToken(token)
	if validErr != nil {
		c.AbortWithStatusJSON(validErr.Status, validErr)
		return
	}

	var user models.User
	user.ID = userID
	err2 := user.FindWithId()
	if err2 == nil {
		auth.SetAuthUser(user)
	}

	c.Next()

}
