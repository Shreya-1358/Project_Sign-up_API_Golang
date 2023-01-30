package service

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"repos/project/model"
	"testing"
)

func TestGetUser(t *testing.T) {

	t.Run("TestGetUser_Success", func(t *testing.T) {

		userList := []*model.User{
			{Id: "int1",
				Name:     "Aishi",
				MobileNo: "8978290148",
				Address:  "Kolkata",
				Username: "Aishi1",
				Password: "start@123"},
			{Id: "int2",
				Name:     "Bali",
				MobileNo: "7820398265",
				Address:  "Patna",
				Username: "Bali2",
				Password: "start@234"},
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().GetAllUsers().Times(1).Return(userList, nil)

		actualUserlist, err := userService.GetAllUsers()
		assert.Equal(t, userList, actualUserlist)
		assert.Nil(t, err)
	})
	t.Run("TestGetUser_ErrorWhileGetting_AllUser", func(t *testing.T) {

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().GetAllUsers().Times(1).Return(nil, errors.New(""))

		actualUserlist, err := userService.GetAllUsers()
		assert.Nil(t, actualUserlist)
		assert.Error(t, err)
	})

}

func TestCreateUser(t *testing.T) {

	t.Run("TestCreateUser_Success", func(t *testing.T) {
		user := model.User{
			Id:       "int3",
			Name:     "Camile",
			MobileNo: "9036790267",
			Address:  "Mumbai",
			Username: "Camile3",
			Password: "start@345",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().CreateUser(&user).Times(1).Return(nil)

		err := userService.CreateUser(&user)
		assert.Nil(t, err)

	})

	t.Run("TestCreateUser_ErrorWhileCreating_NewUser", func(t *testing.T) {
		user := model.User{
			Id:       "int3",
			Name:     "Camile",
			MobileNo: "9036790267",
			Address:  "Mumbai",
			Username: "Camile3",
			Password: "start@345",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().CreateUser(&user).Times(1).Return(errors.New(""))

		err := userService.CreateUser(&user)
		assert.Error(t, err)
	})
}

func TestGetUserById(t *testing.T) {

	t.Run("TestGetUserById_Success", func(t *testing.T) {

		var user *model.User = &model.User{

			Id:       "int4",
			Name:     "Dev",
			MobileNo: "8934567822",
			Address:  "Karnataka",
			Username: "Dev4",
			Password: "start@456",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)
		mockclient.EXPECT().GetUserById(user.Id).Times(1).Return(user, nil)

		actualUser, err := userService.GetUserById(user.Id)
		assert.Equal(t, user, actualUser)
		assert.Nil(t, err)
	})
	t.Run("TestGetUserById_ErrorWhileGetting_User", func(t *testing.T) {
		var user *model.User = &model.User{

			Id:       "int4",
			Name:     "Dev",
			MobileNo: "8934567822",
			Address:  "Karnataka",
			Username: "Dev4",
			Password: "start@456",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)
		mockclient.EXPECT().GetUserById(user.Id).Times(1).Return(nil, errors.New(""))

		actualUser, err := userService.GetUserById(user.Id)
		assert.Nil(t, actualUser)
		assert.Error(t, err)

	})

}

func TestUpdateUser(t *testing.T) {
	t.Run("TestUpdateUser_Success", func(t *testing.T) {
		var user *model.User = &model.User{

			Id:       "int5",
			Name:     "Eden",
			MobileNo: "9278309812",
			Address:  "Tripura",
			Username: "Eden5",
			Password: "start@567",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().GetUserById(user.Id).Times(1).Return(user, nil)
		mockclient.EXPECT().UpdateUser(user, user.Id).Times(1).Return(nil)

		err := userService.UpdateUser(user, user.Id)
		assert.Nil(t, err)
	})

	t.Run("TestUpdateUser_ErrorWhileGetting_User", func(t *testing.T) {

		var user *model.User = &model.User{

			Id:       "int5",
			Name:     "Eden",
			MobileNo: "9278309812",
			Address:  "Tripura",
			Username: "Eden5",
			Password: "start@567",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().GetUserById(user.Id).Times(1).Return(user, errors.New(""))
		mockclient.EXPECT().UpdateUser(user, user.Id).Times(0).Return(nil)

		err := userService.UpdateUser(user, user.Id)
		assert.Error(t, err)
	})
	t.Run("TestUpdateUser_ErrorWhileUpdatingUser", func(t *testing.T) {
		var user *model.User = &model.User{

			Id:       "int5",
			Name:     "Eden",
			MobileNo: "9278309812",
			Address:  "Tripura",
			Username: "Eden5",
			Password: "start@567",
		}

		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserDao(cntrl)
		userService := NewUserServiceImpl(mockclient)

		mockclient.EXPECT().GetUserById(user.Id).Times(1).Return(user, nil)
		mockclient.EXPECT().UpdateUser(user, user.Id).Times(1).Return(errors.New(""))

		err := userService.UpdateUser(user, user.Id)
		assert.Error(t, err)
	})

}
