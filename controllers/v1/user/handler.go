package user

import "github.com/BearTS/go-gin-monolith/services/usersvc"

type userHandler struct {
	usersvc usersvc.Interface
}

func Handler(userSvc usersvc.Interface) *userHandler {
	return &userHandler{
		usersvc: userSvc,
	}
}
