package controllers

type ControllerHub struct {
	User *UserController
}

func NewControllerHub() *ControllerHub {
	return &ControllerHub{
		User: &UserController{},
	}
}