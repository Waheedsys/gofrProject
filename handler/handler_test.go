package handler_test

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"

	gofrHttp "gofr.dev/pkg/gofr/http"
	"gofrProject/entities"
	"gofrProject/handler"
)

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := handler.NewMockUserService(ctrl)
	h := handler.NewUserHandler(mockService)

	tests := []struct {
		name             string
		pathParam        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name:      "successful getuser",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().GetUsers(gomock.Any()).Return([]entities.Users{
					{UserName: "waheed",
						UserAge:     19,
						PhoneNumber: "12345354",
						Email:       "waheed@example.com"},
				}, nil)
			},
			expectedResponse: []entities.Users{
				{UserName: "waheed",
					UserAge:     19,
					PhoneNumber: "12345354",
					Email:       "waheed@example.com"},
			},
			expectedErr: nil,
		},

		{
			name:      "user not found",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().GetUsers(gomock.Any()).Return(nil, errors.New("user not found"))
			},
			expectedResponse: nil,
			expectedErr:      errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user", nil)
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrHttp.NewRequest(gofrHttp.SetPathParam(req, map[string]string{"name": test.pathParam}))

			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			test.mockExpect()

			res, err := h.GetUsers(c)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResponse, res)
		})
	}
}

func Test_GetUserByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := handler.NewMockUserService(ctrl)
	h := handler.NewUserHandler(mockService)

	tests := []struct {
		name             string
		pathParam        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name:      "successful get user by name",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().GetUsersByName("waheed", gomock.Any()).Return(entities.Users{
					UserName:    "waheed",
					UserAge:     19,
					PhoneNumber: "12345354",
					Email:       "waheed@example.com",
				}, nil)
			},
			expectedResponse: entities.Users{
				UserName:    "waheed",
				UserAge:     19,
				PhoneNumber: "12345354",
				Email:       "waheed@example.com",
			},
			expectedErr: nil,
		},
		{
			name:      "user not found",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().GetUsersByName("waheed", gomock.Any()).Return(entities.Users{}, errors.New("user not found"))
			},
			expectedResponse: entities.Users{},
			expectedErr:      errors.New("user not found"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/user/{name}", nil)
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrHttp.NewRequest(gofrHttp.SetPathParam(req, map[string]string{"name": test.pathParam}))
			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			test.mockExpect()

			res, err := h.GetUserByName(c)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResponse, res)
		})
	}
}

func Test_AddUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := handler.NewMockUserService(ctrl)
	h := handler.NewUserHandler(mockService)

	tests := []struct {
		name             string
		inputBody        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name:      "successful add user",
			inputBody: `{"UserName": "waheed", "UserAge": 19, "PhoneNumber": "12345354", "Email": "waheed@example.com"}`,
			mockExpect: func() {
				mockService.EXPECT().AddUsers(gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResponse: nil,
			expectedErr:      nil,
		},

		{
			name:      "error while adding user",
			inputBody: `{"UserName": "waheed", "UserAge": 19, "PhoneNumber": "12345354", "Email": "waheed@example.com"}`,
			mockExpect: func() {
				mockService.EXPECT().AddUsers(gomock.Any(), gomock.Any()).Return(fmt.Errorf("error while adding user"))
			},
			expectedResponse: nil,
			expectedErr:      fmt.Errorf("error while adding user"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/user", strings.NewReader(test.inputBody))
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrHttp.NewRequest(gofrHttp.SetPathParam(req, map[string]string{}))
			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			test.mockExpect()

			res, err := h.AddUser(c)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResponse, res)
		})
	}
}

func Test_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := handler.NewMockUserService(ctrl)
	h := handler.NewUserHandler(mockService)

	tests := []struct {
		name             string
		pathParam        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name:      "successful update user",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().UpdateUsers("waheed", gomock.Any(), gomock.Any()).Return(nil)
			},
			expectedResponse: nil,
			expectedErr:      nil,
		},
		{
			name:      "error while updating user",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().UpdateUsers("waheed", gomock.Any(), gomock.Any()).Return(fmt.Errorf("error while updating user"))
			},
			expectedResponse: nil,
			expectedErr:      fmt.Errorf("error while updating user"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/user/{name}", nil)
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrHttp.NewRequest(gofrHttp.SetPathParam(req, map[string]string{"name": test.pathParam}))
			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			test.mockExpect()

			res, err := h.UpdateUser(c)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResponse, res)
		})
	}
}

func Test_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := handler.NewMockUserService(ctrl)
	h := handler.NewUserHandler(mockService)

	tests := []struct {
		name             string
		pathParam        string
		mockExpect       func()
		expectedResponse interface{}
		expectedErr      error
	}{
		{
			name:      "successful delete user",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().DeleteUsers("waheed", gomock.Any()).Return(nil)
			},
			expectedResponse: nil,
			expectedErr:      nil,
		},
		{
			name:      "error while deleting user",
			pathParam: "waheed",
			mockExpect: func() {
				mockService.EXPECT().DeleteUsers("waheed", gomock.Any()).Return(fmt.Errorf("error while deleting user"))
			},
			expectedResponse: nil,
			expectedErr:      fmt.Errorf("error while deleting user"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodDelete, "/user/{name}", nil)
			req.Header.Set("Content-Type", "application/json")

			gofrR := gofrHttp.NewRequest(gofrHttp.SetPathParam(req, map[string]string{"name": test.pathParam}))
			c := &gofr.Context{
				Context: nil,
				Request: gofrR,
			}
			test.mockExpect()

			res, err := h.DeleteUser(c)

			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, test.expectedResponse, res)
		})
	}
}
