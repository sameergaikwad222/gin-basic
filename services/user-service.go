package services

import (
	"gin-basic/models"
	"gin-basic/repository"
	"gin-basic/utility"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func (s *UserService) GetUsers(queryOption *utility.UserQueryParams) (users []models.User, err error) {
	return s.UserRepo.GetAll(queryOption)
}

func (s *UserService) GetUserById(Id uint) (*models.User, error) {
	return s.UserRepo.GetByID(Id)
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.Create(user)
}

func (s *UserService) CreateBulkUsers(users []*models.User) error {
	return s.UserRepo.CreateBulk(users)
}

func (s *UserService) UpdateUser(id int, user *models.User) error {
	return s.UserRepo.Update(id, user)
}

// func (s *UserService) SoftDeleteUser(id uint) error {
// 	return s.UpdateUser(int(id), *models.User{
// 		gorm.DeletedAt: time.Now(),
// 	})
// }

func (s *UserService) DeleteUser(Id uint) error {
	return s.UserRepo.Delete(Id)
}
