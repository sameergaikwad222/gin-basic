package services

import "gin-basic/repository"

type AllServiceRegistry struct {
	User *UserService
}

func NewServiceRegistry(allRepo *repository.AllRepositoryRegistry) *AllServiceRegistry {
	return &AllServiceRegistry{
		User: &UserService{
			UserRepo: allRepo.UserRepo,
		},
	}
}