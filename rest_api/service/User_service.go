package service

import (
	"repos/project/dao"
	"repos/project/model"
)

type UserService interface {
	GetAllUsers() ([]*model.User, error)
	CreateUser(user *model.User) error
	GetUserById(id string) (*model.User, error)
	UpdateUser(user *model.User, id string) error
}

type UserServiceImpl struct {
	userdao dao.UserDao
}

func NewUserServiceImpl(dao dao.UserDao) UserService {
	return &UserServiceImpl{userdao: dao}
}

func (us *UserServiceImpl) GetAllUsers() ([]*model.User, error) {

	user, err := us.userdao.GetAllUsers()
	if err != nil {
		return user, err
	}
	return user, nil

}

func (us *UserServiceImpl) CreateUser(user *model.User) (err error) {
	err = us.userdao.CreateUser(user)
	if err != nil {
		return err

	}
	return nil

}

func (us *UserServiceImpl) GetUserById(id string) (*model.User, error) {

	user, err := us.userdao.GetUserById(id)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (us *UserServiceImpl) UpdateUser(user *model.User, id string) error {
	_, err := us.userdao.GetUserById(id)
	if err != nil {
		return err
	}
	err = us.userdao.UpdateUser(user, id)
	if err != nil {
		return err
	}

	return nil

}
