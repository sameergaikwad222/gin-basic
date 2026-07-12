package repository

import (
	"gin-basic/models"
	"gin-basic/utility"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}


func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) CreateBulk(users []*models.User) error {
	return r.db.CreateInBatches(users, 10).Error
}

func (r *UserRepository) GetAll(query *utility.UserQueryParams) ([]models.User, error) {
	var users []models.User
	
	// Filter by IDs
	if len(query.IdList) > 0 {
		r.db = r.db.Where("id IN ?", query.IdList)
	}

	// Filter by email
	if query.Email != "" {
		r.db = r.db.Where("email = ?", query.Email)
	}

	// Filter by username
	if query.Username != "" {
		r.db = r.db.Where("username ILIKE ?", "%"+query.Username+"%")
	}

	// Pagination
	if query.Page <= 0 {
		query.Page = 1
	}

	if query.Limit <= 0 {
		query.Limit = 10
	}

	offset := (query.Page - 1) * query.Limit
	r.db = r.db.Offset(offset).Limit(query.Limit)
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error{
	return r.db.Save(user).Error
}	

func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

