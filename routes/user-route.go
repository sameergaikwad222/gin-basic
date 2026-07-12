package routes

import (
	"gin-basic/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, User *controllers.UserController){
	{
		userRouter := rg.Group("/user")
		userRouter.GET("/getUsers", User.GetUsersController)
		userRouter.GET("/:id", User.GetUserByIdController)
		userRouter.POST("/bulkCreate", User.CreateBulkUsersController)
		userRouter.POST(("/create"))
		userRouter.PATCH("/:id", User.UpdateUserController)
		userRouter.DELETE("/:id", User.DeleteUserController)
	}
}