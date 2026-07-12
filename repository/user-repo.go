package repository

import (
	"fmt"
	"gin-basic/models"
	"gin-basic/utility"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) Create(user *models.User) error {
	db := r.db.Model(&models.User{})
	return db.Create(user).Error
}

func (r *UserRepository) CreateBulk(users []*models.User) error {
	db := r.db.Model(&models.User{})
	return db.CreateInBatches(users, 10).Error
}

func (r *UserRepository) GetAll(query *utility.UserQueryParams) ([]models.User, error) {
	var users []models.User
	db := r.db.Model(&models.User{})

	// Filter by IDs
	if len(query.IdList) > 0 {
		db = db.Where("id IN ?", query.IdList)
	}

	// Filter by email
	if query.Email != "" {
		db = db.Where("email = ?", query.Email)
	}

	// Filter by username
	if query.Username != "" {
		db = db.Where("username ILIKE ?", "%"+query.Username+"%")
	}

	// Pagination
	if query.Page <= 0 {
		query.Page = 1
	}

	if query.Limit <= 0 {
		query.Limit = 10
	}

	offset := (query.Page - 1) * query.Limit
	db = db.Offset(offset).Limit(query.Limit)
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	db := r.db.Model(&models.User{})
	err := db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(id int, user *models.User) error {
	updates := map[string]interface{}{}
	db := r.db.Model(&models.User{})
	fmt.Println(user)
	if user.Name != "" {
		updates["Name"] = user.Name
	}

	if user.Age > 0 {
		updates["Age"] = user.Age
	}

	if user.Email != "" {
		updates["Email"] = user.Email
	}

	return db.Model(&models.User{}).Where("id = ?", id).Updates(updates).Error

}

func (r *UserRepository) Delete(id uint) error {
	db := r.db.Model(&models.User{})
	return db.Delete(&models.User{}, id).Error
}
