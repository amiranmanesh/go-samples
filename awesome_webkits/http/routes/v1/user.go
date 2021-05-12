package v1

import (
	auth_controller "awesome_webkits/http/controllers/v1/auth"
	"github.com/gin-gonic/gin"
)

type UserRoute struct{}

func (UserRoute) GetRoutes(router gin.IRouter) {

	//v1
	authRoutes := router.Group("/auth")
	{

		authRoutes.POST("/signup", auth_controller.AuthController.Signup)
		authRoutes.POST("/login", auth_controller.AuthController.Login)

	}

}
