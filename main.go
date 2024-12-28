package main

import (
	"gofr.dev/pkg/gofr"
	"gofrProject/handler"
	"gofrProject/service"
	"gofrProject/store"
)

func main() {
	// Create a new application
	a := gofr.New()
	userstore := store.NewDetails()
	userService := service.NewUserService(userstore)
	userHandler := handler.NewUserHandler(userService)

	a.GET("/user", userHandler.GetUsers)
	a.POST("/user", userHandler.AddUser)
	a.GET("/user/{name}", userHandler.GetUserByName)
	a.PUT("/user/{name}", userHandler.UpdateUser)
	a.DELETE("/user/{name}", userHandler.DeleteUser)
	a.UseMiddleware(Authentication)
	a.Run()
}
