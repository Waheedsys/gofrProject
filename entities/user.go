package entities

import (
	"errors"
)

var ErrInvalidPhoneNumber = errors.New("invalid phone number")

type Users struct {
	UserName    string `json:"user_name"`
	UserAge     int    `json:"user_age"`
	PhoneNumber string `json:"phone_Number"`
	Email       string `json:"email"`
}
