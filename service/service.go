package service

import (
	"database/sql"
	"fmt"
	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/datasource"
	"gofr.dev/pkg/gofr/http"
	"gofrProject/entities"
)

type Service struct {
	store UserStore
}

func NewUserService(store UserStore) *Service {
	return &Service{store: store}
}

func (s *Service) GetUsers(ctx *gofr.Context) ([]entities.Users, error) {
	users, err := s.store.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *Service) GetUsersByName(name string, ctx *gofr.Context) (entities.Users, error) {
	user, err := s.store.GetUsersByName(name, ctx)
	if err != nil {
		if err == sql.ErrNoRows {

			return entities.Users{}, fmt.Errorf("%w, '%s'", http.ErrorEntityNotFound{"name", "name"}, "john")
		}

		return entities.Users{}, err
	}

	return user, nil
}

func (s *Service) AddUsers(user *entities.Users, ctx *gofr.Context) error {

	if user.UserName == "" || user.PhoneNumber == "" {
		dbErr2 := datasource.ErrorDB{Message: "UserName and PhoneNumber cannot be empty"}
		return fmt.Errorf("error:%w", dbErr2)
	}

	existingUser, err := s.store.GetUsersByName(user.UserName, ctx)
	if err == nil && existingUser.UserName != "" {
		var err1 = http.ErrorEntityAlreadyExist{}
		return fmt.Errorf("%w, '%s' already exists", err1, user.UserName)
	}

	return s.store.AddUsers(user, ctx)
}

func (s *Service) DeleteUsers(name string, ctx *gofr.Context) error {
	existingUser, err := s.store.GetUsersByName(name, ctx)
	if err != nil || existingUser.UserName == "" {
		return fmt.Errorf("%w", http.ErrorEntityNotFound{"name", "albert"})
	}

	return s.store.DeleteUsers(name, ctx)
}

func (s *Service) UpdateUsers(name string, updateUser *entities.Users, ctx *gofr.Context) error {
	existingUser, err := s.store.GetUsersByName(name, ctx)
	if err != nil || existingUser.UserName == "" {
		return fmt.Errorf("%w", http.ErrorEntityNotFound{"name", "albert"})
	}

	return s.store.UpdateUsers(name, updateUser, ctx)
}
