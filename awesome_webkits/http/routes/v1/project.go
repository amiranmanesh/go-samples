package v1

import (
	"awesome_webkits/http/controllers/v1/project"
	"awesome_webkits/http/middleware"
	"github.com/gin-gonic/gin"
)

type ProjectRoute struct{}

func (ProjectRoute) GetRoutes(router gin.IRouter) {

	//v1/project
	authRoutes := router.Group("/project")
	{

		authRoutes.Use(middleware.RestAuthUserMiddleware)
		authRoutes.POST("/create", project.ProjectController.Save)
		authRoutes.POST("/edit", project.ProjectController.NameEdit)
		authRoutes.POST("/api/create", project.ProjectApiController.Save)
		authRoutes.POST("/api/edit", project.ProjectApiController.Edit)
	}

	apiRoutes := router.Group("/project")
	{
		//no need to auth
		apiRoutes.GET("/api/:path", project.ProjectApiController.Show)
	}

}
