package v1

import (
	auth "awesome_webkits/http/controllers/v1/activation"
	"github.com/gin-gonic/gin"
)

type ActivationRoute struct{}

func (ActivationRoute) GetRoutes(router gin.IRouter) {

	//v1/activation
	apiRoutes := router.Group("/activation")
	{
		//no need to auth
		apiRoutes.GET("/email/:token", auth.ActivationController.Email)
		apiRoutes.GET("/pass/:token", auth.ActivationController.Password)
	}

}
