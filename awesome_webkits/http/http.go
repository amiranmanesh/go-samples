package http

import (
	"awesome_webkits/http/routes"
	v1rotues "awesome_webkits/http/routes/v1"
	"awesome_webkits/utils/env"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

type iHttp interface {
	Initialize()
	Run()
}

type restHttp struct{}

var App iHttp = &restHttp{}

var (
	router *gin.Engine
)

func (restHttp) Initialize() {
	router = gin.Default()
	// a.Router.SetMode(gin.ReleaseMode)
	setRouters()
}

func setRouters() {
	routes.SwaggerRoute{}.GetRoutes(router)
	routes.GraphqlRoute{}.GetRoutes(router)
	v1 := router.Group("/v1")
	{
		v1rotues.UserRoute{}.GetRoutes(v1)
		v1rotues.ProjectRoute{}.GetRoutes(v1)
		v1rotues.ActivationRoute{}.GetRoutes(v1)
		v1rotues.UploadRoute{}.GetRoutes(v1)
	}
}

//run.sh
func (restHttp) Run() {
	host := ":" + env.GetEnvItem("HTTP_PORT")
	logrus.WithFields(logrus.Fields{
		"port": host,
	}).Info("app now is listening and serving...")
	logrus.Error(errors.Trace(router.Run(host)))
}
