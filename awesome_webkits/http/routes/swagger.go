package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"net/http"
)

type SwaggerRoute struct{}

func (SwaggerRoute) GetRoutes(router gin.IRouter) {
	router.GET("/docs", docksHandler())
	router.GET("/swagger.yaml", swaggerHandler())
}

// Defining the Graphql handler
func docksHandler() gin.HandlerFunc {
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	return func(c *gin.Context) {
		sh.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func swaggerHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		http.FileServer(http.Dir("./"))
	}
}
