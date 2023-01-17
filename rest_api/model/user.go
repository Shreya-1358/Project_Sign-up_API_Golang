package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	Config "repos/project/config"
)

// GetAllUsers Fetch all user data
func GetAllUsers(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser ... Insert New data
func CreateUser(user *User) (err error) {
	fmt.Println("user:", *user)
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

// GetUserByID ... Fetch only one user by Id
func GetUserByID(user *User, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

// UpdateUser ... Update user
func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}

// DeleteUser ... Delete user
func DeleteUser(user *User, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}

//type NewInterface interface {
//	getName()
//}
//
//type user struct{}
//
//func (s *user) name() string {
//	return "abc"
//}
