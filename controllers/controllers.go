package controllers

import (
	"gin-basic/services"
)
type ControllerHub struct {
	User *UserController
}

func NewControllerHub(registry *services.AllServiceRegistry) *ControllerHub {
	return &ControllerHub{
		User: &UserController{
			userService: registry.User,
		},
	}
}