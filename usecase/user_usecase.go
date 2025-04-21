package usecase

import (
	"github.com/witchakornb/basic-ecommerce/domain/entity"
	"github.com/witchakornb/basic-ecommerce/domain/repository"
)

type UserUseCase interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUserByID(id int) (entity.User, error)
	UpdateUser(user entity.User) (entity.User, error)
	DeleteUser(id int) error
}

type UserUseCaseImpl struct {
	UserRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{
		UserRepo: userRepo,
	}
}

func (u *UserUseCaseImpl) CreateUser(user entity.User) (entity.User, error) {
	user, err := u.UserRepo.CreateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCaseImpl) GetUserByID(id int) (entity.User, error) {
	user, err := u.UserRepo.GetUserByID(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCaseImpl) UpdateUser(user entity.User) (entity.User, error) {
	user, err := u.UserRepo.UpdateUser(user)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *UserUseCaseImpl) DeleteUser(id int) error {
	err := u.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
