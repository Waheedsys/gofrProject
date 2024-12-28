package store

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
	"gofrProject/entities"
	"log"
)

// UsersList is a struct that represents the user store with a connection to the database.
type UsersList struct {
	db *sql.DB
}

// NewDetails creates a new instance of UsersList.
func NewDetails() *UsersList {
	return &UsersList{}
}

// GetUsers retrieves all users from the database.
func (userStore *UsersList) GetUsers(ctx *gofr.Context) ([]entities.Users, error) {
	// Query the database for all users.
	rows, err := ctx.SQL.Query("SELECT UserName, UserAge, PhoneNumber, Email FROM User")
	if err != nil {
		// Return a custom error if the SQL query fails.
		dbErr := datasource.ErrorDB{Err: fmt.Errorf("some db error"), Message: "error from sql db"}
		return nil, dbErr
	}
	defer rows.Close()

	var users []entities.Users
	// Iterate through the rows and scan the user details into the struct.
	for rows.Next() {
		var user entities.Users
		if err := rows.Scan(&user.UserName, &user.UserAge, &user.PhoneNumber, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	// Return nil if no users are found.
	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

// GetUsersByName retrieves a single user by their username.
func (userStore *UsersList) GetUsersByName(name string, ctx *gofr.Context) (entities.Users, error) {
	var user entities.Users
	// Query the database for a user by their username.
	err := ctx.SQL.QueryRow("SELECT UserName, UserAge, PhoneNumber, Email FROM User WHERE Username = ?", name).
		Scan(&user.UserName, &user.UserAge, &user.PhoneNumber, &user.Email)
	// If no user is found, return an error.
	if errors.Is(err, sql.ErrNoRows) {
		return entities.Users{}, fmt.Errorf("user with name '%v'not found", name)
	}
	// Return the user if found.
	return user, nil
}

// AddUsers inserts a new user into the database.
func (userStore *UsersList) AddUsers(user *entities.Users, ctx *gofr.Context) error {
	log.Printf("Inserting user: %+v", user)
	// Check if UserName or PhoneNumber is empty
	if user.UserName == "" || user.PhoneNumber == "" {
		return fmt.Errorf("UserName and PhoneNumber cannot be empty")
	}
	//  Exec the database for addding the user .
	_, err := ctx.SQL.Exec("INSERT INTO User (UserName, UserAge, PhoneNumber, Email) VALUES (?, ?, ?, ?)",
		user.UserName, user.UserAge, user.PhoneNumber, user.Email)
	// If unable to add user, return error
	if err != nil {
		dbErr := datasource.ErrorDB{Err: err, Message: "error from sql db"}
		return dbErr
	}

	return nil
}

// DeleteUsers a user from the database.
func (userStore *UsersList) DeleteUsers(name string, ctx *gofr.Context) error {
	_, err := ctx.SQL.Exec("DELETE FROM User WHERE UserName = ?", name)
	return err
}

// UpdateUsers a user from the database.
func (userStore *UsersList) UpdateUsers(name string, updateUser *entities.Users, ctx *gofr.Context) error {
	_, err := ctx.SQL.Exec("UPDATE User SET Email = ? WHERE UserName = ?", updateUser.Email, name)
	return err
}
