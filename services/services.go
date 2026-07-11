package services

type ServiceHub struct {
	User *UserService
}

func NewServiceHub() *ServiceHub {
	return &ServiceHub{
		User: &UserService{},
	}
}