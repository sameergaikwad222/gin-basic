package controllers

import (
	"gin-basic/models"
	"gin-basic/services"
	"gin-basic/utility"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func (uc *UserController) GetUsersController(c *gin.Context) {
	queryParams := &utility.UserQueryParams{}
	if err := c.ShouldBindQuery(&queryParams); err != nil {
		sendErrorResponse(400,err,c)
		return
	}

	users, err := uc.userService.GetUsers(queryParams)
	if err != nil {
		sendErrorResponse(500,err,c)
		return
	}
	c.JSON(200, gin.H{"users": users})
}

func (uc *UserController) GetUserByIdController(c *gin.Context) {
	Idstr := c.Param("id")
	id,err := strconv.ParseUint(Idstr,10,64)
	if err != nil {
		sendErrorResponse(400, err, c)
		return 
	}
	user, err := uc.userService.GetUserById(uint(id))
	if err != nil{
		sendErrorResponse(500, err, c)
		return 
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateBulkUsersController(c *gin.Context) {
	var users [] *models.User
	if err := c.ShouldBindBodyWithJSON(users); err != nil {
		sendErrorResponse(400,err,c)
		return
	}
	if err1 := uc.userService.CreateBulkUsers(users); err1 != nil {
		sendErrorResponse(500,err1,c)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Bulk Users Inserted!",
	})
}

func (uc *UserController) CreateUserController(c *gin.Context){
	var user *models.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		sendErrorResponse(400,err,c)
		return
	}
	err1 := uc.userService.CreateUser(user)
	if err1 != nil {
		sendErrorResponse(500,err1,c)
		return
	}
	c.JSON(http.StatusCreated, gin.H {
		"message" : "User Created Successfully",
	})
}

func (uc *UserController) UpdateUserController(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		sendErrorResponse(400, err, c)
	}
	err1 := uc.userService.UpdateUser(user)
	if err1 != nil {
		sendErrorResponse(500, err1, c)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "user updated successfully",
	})

}

func (uc *UserController) DeleteUserController(c *gin.Context) {
	Idstr := c.Param("id")
	id,err := strconv.ParseUint(Idstr,10,64)
	if err != nil {
		sendErrorResponse(400, err, c)
		return 
	}
	err1 := uc.userService.DeleteUser(uint(id))
	if err1 != nil{
		sendErrorResponse(500, err1, c)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message" : "user deleted successfully",
	})
}

func sendErrorResponse(statusCode uint,err error, c *gin.Context){
	switch statusCode {
	case 400:
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err,
		})
	case 500:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error" : err,
		})
	}
}