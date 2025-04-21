package infrastructure

import (
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &gormUserRepository{db: db}
}

// CreateUser creates a new user in the database.
func (r *gormUserRepository) CreateUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// GetUserByID retrieves a user by ID from the database.
func (r *gormUserRepository) GetUserByID(id int) (entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// GetAllUsers retrieves all users from the database.
func (r *gormUserRepository) UpdateUser(user entity.User) (entity.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

// DeleteUser deletes a user by ID from the database.
func (r *gormUserRepository) DeleteUser(id int) error {
	var user entity.User
	err := r.db.Delete(&user, id).Error
	if err != nil {
		return err
	}
	return nil
}
