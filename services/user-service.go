package services

type UserService struct{}

func (s *UserService) GetUsers() ([]models.User, err) {
	return []models.User{}, nil
}