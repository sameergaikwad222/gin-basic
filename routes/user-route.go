package routes

import (
	"gin-basic/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, ctrlHub *controllers.ControllerHub){
	{
		userRouter := rg.Group("/user")
		userRouter.GET("/getUsers", ctrlHub.User.GetUsersController)
		userRouter.POST("/createUsers", ctrlHub.User.CreateBulkUsersController)
		userRouter.PATCH("/:id", ctrlHub.User.UpdateUserController)
		userRouter.DELETE("/:id", ctrlHub.User.DeleteUserController)
	}
}