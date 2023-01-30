package dao

import (
	Config "repos/project/config"
	"repos/project/model"
)

type UserDao interface {
	GetAllUsers() ([]*model.User, error)
	CreateUser(user *model.User) error
	GetUserById(id string) (*model.User, error)
	UpdateUser(user *model.User, id string) error
}

type UserDaoImpl struct {
}

func NewUserDaoImpl() *UserDaoImpl {
	return &UserDaoImpl{}
}

func (ud *UserDaoImpl) GetAllUsers() ([]*model.User, error) {
	var user []*model.User
	if err := Config.DB.Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil

}

func (ud *UserDaoImpl) CreateUser(user *model.User) (err error) {

	if err = Config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ud *UserDaoImpl) GetUserById(id string) (*model.User, error) {
	var user *model.User
	if err := Config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (ud *UserDaoImpl) UpdateUser(user *model.User, id string) (err error) {
	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}
