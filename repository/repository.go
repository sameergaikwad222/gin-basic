package repository

import "gorm.io/gorm"

type AllRepositoryRegistry struct {
	UserRepo *UserRepository
}

func NewRepositoryRegistry(db *gorm.DB) *AllRepositoryRegistry {
	return &AllRepositoryRegistry{
		UserRepo: &UserRepository{
			db: db,
		},
	}
}