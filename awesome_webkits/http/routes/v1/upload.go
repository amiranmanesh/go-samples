package v1

import (
	"awesome_webkits/http/controllers/v1/file"
	"awesome_webkits/http/middleware"
	"github.com/gin-gonic/gin"
)

type UploadRoute struct{}

func (UploadRoute) GetRoutes(router gin.IRouter) {
	//v1/upload
	router.Use(middleware.RestAuthUserMiddleware)
	router.POST("/upload", file.UploadController.Save)

}
