package auth

import (
	"awesome_webkits/database/models"
)

var authUser models.User

/**
Singleton auth_user
*/
func GetAuthUser() models.User {
	return authUser
}

func SetAuthUser(user models.User) {
	authUser = user
}
