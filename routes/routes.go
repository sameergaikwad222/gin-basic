package routes

import (
	"gin-basic/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAllRoutes(router *gin.Engine){

	v1 := router.Group("/api/v1")
	ctrlHub := controllers.NewControllerHub()

	RegisterUserRoutes(v1, ctrlHub)
	// RegisterProductRoutes(v1)
}