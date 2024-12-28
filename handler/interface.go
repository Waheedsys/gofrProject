package handler

import (
	"gofr.dev/pkg/gofr"
	"gofrProject/entities"
)

type UserService interface {
	GetUsers(ctx *gofr.Context) ([]entities.Users, error)
	GetUsersByName(name string, ctx *gofr.Context) (entities.Users, error)
	AddUsers(user *entities.Users, ctx *gofr.Context) error
	DeleteUsers(name string, ctx *gofr.Context) error
	UpdateUsers(name string, updateUser *entities.Users, ctx *gofr.Context) error
}
