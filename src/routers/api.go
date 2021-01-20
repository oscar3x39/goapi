package routers

import (
	"goapi/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter is gin routers
func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/books", controllers.ListBook)
	}

	return r
}