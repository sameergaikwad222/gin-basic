package routes

import (
	"gin-basic/controllers"
	"gin-basic/repository"
	"gin-basic/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAllRoutes(router *gin.Engine, db *gorm.DB){

	v1 := router.Group("/api/v1")
	reporegistry := repository.NewRepositoryRegistry(db)
	serviceRegistry := services.NewServiceRegistry(reporegistry)
	ctrlHub := controllers.NewControllerHub(serviceRegistry)

	RegisterUserRoutes(v1, ctrlHub.User)
	// RegisterProductRoutes(v1, ctrlHub.Product)
}