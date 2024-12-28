package service

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/http"
	"gofrProject/entities"
	"testing"
)

func Test_GetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	users := []entities.Users{
		{UserName: "john", PhoneNumber: "1234"},
		{UserName: "alice", PhoneNumber: "5678"},
	}

	mockStore.EXPECT().GetUsers(gomock.Any()).Return(users, nil).Times(1)

	ctx := &gofr.Context{}
	result, err := service.GetUsers(ctx)

	assert.NoError(t, err)
	assert.Equal(t, users, result)
}

func Test_GetUsersByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	tests := []struct {
		name        string
		mockReturn  entities.Users
		mockError   error
		expected    entities.Users
		expectedErr error
	}{
		{
			name:        "Existing User",
			mockReturn:  entities.Users{UserName: "john", PhoneNumber: "1234"},
			mockError:   nil,
			expected:    entities.Users{UserName: "john", PhoneNumber: "1234"},
			expectedErr: nil,
		},
		{
			name:        "User Not Found",
			mockReturn:  entities.Users{},
			mockError:   sql.ErrNoRows,
			expected:    entities.Users{},
			expectedErr: fmt.Errorf("%w, '%s'", http.ErrorEntityNotFound{"name", "name"}, "john"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockStore.EXPECT().GetUsersByName(tt.name, gomock.Any()).Return(tt.mockReturn, tt.mockError).Times(1)

			ctx := &gofr.Context{}
			result, err := service.GetUsersByName(tt.name, ctx)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func Test_AddUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	user := &entities.Users{UserName: "john", PhoneNumber: "1234"}

	tests := []struct {
		name        string
		mockReturn  entities.Users
		mockError   error
		expectedErr error
	}{
		{
			name:        "Valid User",
			mockReturn:  entities.Users{},
			mockError:   nil,
			expectedErr: nil,
		},
		{
			name:        "User Already Exists",
			mockReturn:  entities.Users{UserName: "john", PhoneNumber: "1234"},
			mockError:   nil,
			expectedErr: fmt.Errorf("%w, '%s' already exists", http.ErrorEntityAlreadyExist{}, user.UserName),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.name == "User Already Exists" {

				mockStore.EXPECT().GetUsersByName(user.UserName, gomock.Any()).Return(tt.mockReturn, nil).Times(1)

				mockStore.EXPECT().AddUsers(gomock.Any(), gomock.Any()).Times(0)
			} else {

				mockStore.EXPECT().GetUsersByName(user.UserName, gomock.Any()).Return(entities.Users{}, sql.ErrNoRows).Times(1)

				mockStore.EXPECT().AddUsers(user, gomock.Any()).Return(nil).Times(1)
			}

			err := service.AddUsers(user, &gofr.Context{})

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_DeleteUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	tests := []struct {
		name        string
		mockReturn  entities.Users
		mockError   error
		expectedErr error
	}{
		{
			name:        "User Exists",
			mockReturn:  entities.Users{UserName: "john", PhoneNumber: "1234"},
			mockError:   nil,
			expectedErr: nil,
		},
		{
			name:        "User Not Found",
			mockReturn:  entities.Users{},
			mockError:   sql.ErrNoRows,
			expectedErr: fmt.Errorf("No entity found with name: albert"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock behavior for GetUsersByName
			mockStore.EXPECT().GetUsersByName(tt.name, gomock.Any()).Return(tt.mockReturn, tt.mockError).Times(1)

			if tt.expectedErr != nil {
				err := service.DeleteUsers(tt.name, &gofr.Context{})
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				mockStore.EXPECT().DeleteUsers(tt.name, gomock.Any()).Return(nil).Times(1)
				err := service.DeleteUsers(tt.name, &gofr.Context{})
				assert.NoError(t, err)
			}
		})
	}
}

func Test_UpdateUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	tests := []struct {
		name        string
		mockReturn  entities.Users
		mockError   error
		expectedErr error
	}{
		{
			name:        "User Exists",
			mockReturn:  entities.Users{UserName: "john", PhoneNumber: "1234"},
			mockError:   nil,
			expectedErr: nil,
		},
		{
			name:        "User Not Found",
			mockReturn:  entities.Users{},
			mockError:   sql.ErrNoRows,
			expectedErr: fmt.Errorf("No entity found with name: albert"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock behavior for GetUsersByName
			mockStore.EXPECT().GetUsersByName(tt.name, gomock.Any()).Return(tt.mockReturn, tt.mockError).Times(1)

			if tt.expectedErr != nil {
				err := service.UpdateUsers(tt.name, &entities.Users{}, &gofr.Context{})
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				mockStore.EXPECT().UpdateUsers(tt.name, gomock.Any(), gomock.Any()).Return(nil).Times(1)
				err := service.UpdateUsers(tt.name, &entities.Users{}, &gofr.Context{})
				assert.NoError(t, err)
			}
		})
	}
}
