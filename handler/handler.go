package handler

import (
	"fmt"
	"gofr.dev/pkg/gofr"
	"gofrProject/entities"
)

type Handler struct {
	UserService UserService
}

func NewUserHandler(userService UserService) *Handler {
	return &Handler{UserService: userService}
}

func (h *Handler) GetUsers(ctx *gofr.Context) (any, error) {
	resp, err := h.UserService.GetUsers(ctx)
	if err != nil {
		fmt.Sprintf("error while getting user: %s", err)
		return nil, err
	}
	return resp, nil
}

func (h *Handler) GetUserByName(ctx *gofr.Context) (interface{}, error) {
	name := ctx.Request.PathParam("name")

	resp, err := h.UserService.GetUsersByName(name, ctx)
	if err != nil {
		fmt.Sprintf("error while getting user by name: %s", err)
		return resp, err
	}
	return resp, nil
}

func (h *Handler) AddUser(ctx *gofr.Context) (interface{}, error) {
	var newUser entities.Users

	if err := ctx.Bind(&newUser); err != nil {
		return nil, fmt.Errorf("error while adding user: %v", err)
	}

	if err := h.UserService.AddUsers(&newUser, ctx); err != nil {
		return nil, err
	}

	return nil, nil
}

func (h *Handler) UpdateUser(ctx *gofr.Context) (interface{}, error) {

	name := ctx.Request.PathParam("name")

	var updateUser entities.Users

	if err := h.UserService.UpdateUsers(name, &updateUser, ctx); err != nil {
		fmt.Sprintf("error while updating user: %s", err)
		return nil, err
	}
	return nil, nil
}

func (h *Handler) DeleteUser(ctx *gofr.Context) (interface{}, error) {

	name := ctx.Request.PathParam("name")

	if err := h.UserService.DeleteUsers(name, ctx); err != nil {
		fmt.Sprintf("error while deleting user: %s", err)
		return nil, err
	}
	return nil, nil
}
