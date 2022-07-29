package controllers

import (
	"github.com/solrac97gr/go-template/internals/users/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type UserControllers struct {
	app    domain.UserApplication
	logger *logger.Logger
}

func NewUserControllers(app domain.UserApplication, logger *logger.Logger) *UserControllers {
	return &UserControllers{
		app:    app,
		logger: logger,
	}
}
