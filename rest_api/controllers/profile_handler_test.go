package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
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

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)

		mockclient.EXPECT().GetAllUsers().Return(userList, nil).Times(1)

		server := gin.Default()
		server.GET("User", userController.GetUser)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodGet, "/User", nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("TestGetUser_InternalServerError", func(t *testing.T) {
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

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().GetAllUsers().Return(userList, errors.New("")).Times(1)

		server := gin.Default()
		server.GET("User", userController.GetUser)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodGet, "/User", nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)

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

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().CreateUser(&user).Return(nil)

		server := gin.Default()
		data, err := json.Marshal(user)
		require.NoError(t, err)
		server.POST("User", userController.CreateUser)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodPost, "/User", bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)

	})
	t.Run("TestCreateUser_InternalServerError", func(t *testing.T) {
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

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().CreateUser(gomock.Any()).Return(errors.New(""))

		server := gin.Default()
		data, err := json.Marshal(user)
		require.NoError(t, err)
		server.POST("User", userController.CreateUser)
		recorder := httptest.NewRecorder()

		request := httptest.NewRequest(http.MethodPost, "/User", bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)
	})
}

func TestGetUserById(t *testing.T) {

	t.Run("TestGetUserById_Success", func(t *testing.T) {
		user := model.User{
			Id:       "int4",
			Name:     "Dev",
			MobileNo: "8934567822",
			Address:  "Karnataka",
			Username: "Dev4",
			Password: "start@456",
		}
		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().GetUserById(user.Id).Return(&user, nil)

		server := gin.Default()

		server.GET("User/:id", userController.GetUserById)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/User/%s", user.Id)
		request := httptest.NewRequest(http.MethodGet, url, nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("TestGetUserById_InternalServerError", func(t *testing.T) {
		user := model.User{
			Id:       "int4",
			Name:     "Dev",
			MobileNo: "8934567822",
			Address:  "Karnataka",
			Username: "Dev4",
			Password: "start@456",
		}
		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)

		mockclient.EXPECT().GetUserById(user.Id).Return(&user, errors.New(""))

		server := gin.Default()

		server.GET("User/:id", userController.GetUserById)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/User/%s", user.Id)
		request := httptest.NewRequest(http.MethodGet, url, nil)

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)

	})
}

func TestUpdateUser(t *testing.T) {

	t.Run("TestUpdateUser_Success", func(t *testing.T) {
		user := model.User{
			Id:       "int5",
			Name:     "Eden",
			MobileNo: "9278309812",
			Address:  "Tripura",
			Username: "Eden5",
			Password: "start@567",
		}
		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().UpdateUser(&user, user.Id).Times(1).Return(nil)

		server := gin.Default()
		data, err := json.Marshal(user)
		require.NoError(t, err)

		server.PUT("User/:id", userController.UpdateUser)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/User/%s", user.Id)
		request := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("TestUpdateUser_InternalServerError", func(t *testing.T) {
		user := model.User{
			Id:       "int5",
			Name:     "Eden",
			MobileNo: "9278309812",
			Address:  "Tripura",
			Username: "Eden5",
			Password: "start@567",
		}
		cntrl := gomock.NewController(t)
		defer cntrl.Finish()

		mockclient := NewMockUserService(cntrl)
		userController := NewUserController(mockclient)
		mockclient.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(errors.New(""))

		server := gin.Default()
		data, err := json.Marshal(user)
		require.NoError(t, err)
		server.PUT("User/:id", userController.UpdateUser)
		recorder := httptest.NewRecorder()
		url := fmt.Sprintf("/User/%s", user.Id)
		request := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

		server.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusInternalServerError, recorder.Code)

	})
}
