package store

import (
	"database/sql"
	"fmt"
	"gofr.dev/pkg/gofr/datasource"
	"golang.org/x/net/context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/container"
	"gofrProject/entities"
)

func TestGetUsers(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	tests := []struct {
		name             string
		mockExpect       func()
		expectedResponse interface{}
		expectedError    error
	}{
		{
			name: "Successful retrieval of users",
			mockExpect: func() {
				mock.SQL.ExpectQuery("SELECT UserName, UserAge, PhoneNumber, Email FROM User").
					WillReturnRows(sqlmock.NewRows([]string{"UserName", "UserAge", "PhoneNumber", "Email"}).
						AddRow("John Doe", 30, "123-456-7890", "john@example.com"))
			},
			expectedResponse: []entities.Users{
				{
					UserName:    "John Doe",
					UserAge:     30,
					PhoneNumber: "123-456-7890",
					Email:       "john@example.com",
				},
			},
			expectedError: nil,
		},
		{
			name: "Error while fetching users",
			mockExpect: func() {
				mock.SQL.ExpectQuery("SELECT UserName, UserAge, PhoneNumber, Email FROM User").
					WillReturnError(fmt.Errorf("some db error"))
			},
			expectedResponse: []entities.Users([]entities.Users(nil)),
			expectedError: datasource.ErrorDB{
				Err:     fmt.Errorf("some db error"),
				Message: "error from sql db",
			},
		},
		{
			name: "No users found",
			mockExpect: func() {
				mock.SQL.ExpectQuery("SELECT UserName, UserAge, PhoneNumber, Email FROM User").
					WillReturnRows(sqlmock.NewRows([]string{"UserName", "UserAge", "PhoneNumber", "Email"}))
			},
			expectedResponse: []entities.Users([]entities.Users(nil)),
			expectedError:    nil,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpect()

			store := NewDetails()
			users, err := store.GetUsers(ctx)

			assert.Equal(t, tt.expectedResponse, users, "TEST[%d] failed: %s", i, tt.name)
			assert.Equal(t, tt.expectedError, err, "TEST[%d] failed: %s", i, tt.name)
		})
	}
}

func TestGetUsersByName(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	tests := []struct {
		name             string
		username         string
		mockExpect       func()
		expectedResponse interface{}
		expectedError    error
	}{
		{
			name:     "User found",
			username: "John Doe",
			mockExpect: func() {
				mock.SQL.ExpectQuery("SELECT UserName, UserAge, PhoneNumber, Email FROM User WHERE Username = ?").
					WithArgs("John Doe").
					WillReturnRows(sqlmock.NewRows([]string{"UserName", "UserAge", "PhoneNumber", "Email"}).
						AddRow("John Doe", 30, "123-456-7890", "john@example.com"))
			},
			expectedResponse: entities.Users{
				UserName:    "John Doe",
				UserAge:     30,
				PhoneNumber: "123-456-7890",
				Email:       "john@example.com",
			},
			expectedError: nil,
		},
		{
			name:     "User not found",
			username: "Jane Doe",
			mockExpect: func() {
				mock.SQL.ExpectQuery("SELECT UserName, UserAge, PhoneNumber, Email FROM User WHERE Username = ?").
					WithArgs("Jane Doe").
					WillReturnError(sql.ErrNoRows)
			},
			expectedResponse: entities.Users{},
			expectedError:    fmt.Errorf("user with name 'Jane Doe'not found"),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpect()

			store := NewDetails()
			user, err := store.GetUsersByName(tt.username, ctx)

			assert.Equal(t, tt.expectedResponse, user, "TEST[%d] failed: %s", i, tt.name)
			assert.Equal(t, tt.expectedError, err, "TEST[%d] failed: %s", i, tt.name)
		})
	}
}

func TestAddUsers(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	tests := []struct {
		name             string
		user             *entities.Users
		mockExpect       func()
		expectedResponse error
	}{
		{
			name: "Successful user addition",
			user: &entities.Users{
				UserName:    "John Doe",
				UserAge:     30,
				PhoneNumber: "123-456-7890",
				Email:       "john@example.com",
			},
			mockExpect: func() {
				mock.SQL.ExpectExec("INSERT INTO User (UserName, UserAge, PhoneNumber, Email) VALUES (?, ?, ?, ?)").
					WithArgs("John Doe", 30, "123-456-7890", "john@example.com").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedResponse: nil,
		},
		{
			name: "Error on user addition due to empty fields",
			user: &entities.Users{
				UserName:    "",
				UserAge:     30,
				PhoneNumber: "",
				Email:       "john@example.com",
			},
			mockExpect:       func() {},
			expectedResponse: fmt.Errorf("UserName and PhoneNumber cannot be empty"),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpect()

			store := NewDetails()
			err := store.AddUsers(tt.user, ctx)

			assert.Equal(t, tt.expectedResponse, err, "TEST[%d] failed: %s", i, tt.name)
		})
	}
}

func TestDeleteUsers(t *testing.T) {
	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	tests := []struct {
		name             string
		username         string
		mockExpect       func()
		expectedResponse error
	}{
		{
			name:     "Successful deletion",
			username: "John Doe",
			mockExpect: func() {
				mock.SQL.ExpectExec("DELETE FROM User WHERE UserName = ?").
					WithArgs("John Doe").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedResponse: nil,
		},
		{
			name:     "Error while deleting user",
			username: "John Doe",
			mockExpect: func() {
				mock.SQL.ExpectExec("DELETE FROM User WHERE UserName = ?").
					WithArgs("John Doe").
					WillReturnError(fmt.Errorf("db error"))
			},
			expectedResponse: fmt.Errorf("db error"),
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockExpect()

			store := NewDetails()
			err := store.DeleteUsers(tt.username, ctx)

			assert.Equal(t, tt.expectedResponse, err, "TEST[%d] failed: %s", i, tt.name)
		})
	}
}

// test update.
func TestUpdateUsers(t *testing.T) {

	mockContainer, mock := container.NewMockContainer(t)
	ctx := &gofr.Context{
		Context:   context.Background(),
		Container: mockContainer,
	}

	name := "John Doe"
	updateUser := &entities.Users{
		Email: "john.new@example.com",
	}

	tests := []struct {
		name          string
		mockExpect    func()
		expectedError error
	}{
		{
			name: "Successful update",
			mockExpect: func() {

				mock.SQL.ExpectExec("UPDATE User SET Email = ? WHERE UserName = ?").
					WithArgs(updateUser.Email, name).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedError: nil,
		},
		{
			name: "Error while updating user",
			mockExpect: func() {

				mock.SQL.ExpectExec("UPDATE User SET Email = ? WHERE UserName = ?").
					WithArgs(updateUser.Email, name).
					WillReturnError(fmt.Errorf("database error"))
			},
			expectedError: fmt.Errorf("database error"),
		},
		{
			name: "No rows affected",
			mockExpect: func() {

				mock.SQL.ExpectExec("UPDATE User SET Email = ? WHERE UserName = ?").
					WithArgs(updateUser.Email, name).
					WillReturnResult(sqlmock.NewResult(0, 0))
			},
			expectedError: nil,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.mockExpect()

			store := &UsersList{}
			err := store.UpdateUsers(name, updateUser, ctx)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error(), "TEST[%d] failed: %s", i, tt.name)
			} else {
				assert.NoError(t, err, "TEST[%d] failed: %s", i, tt.name)
			}

			assert.NoError(t, mock.SQL.ExpectationsWereMet(), "TEST[%d] failed: %s", i, tt.name)
		})
	}
}
