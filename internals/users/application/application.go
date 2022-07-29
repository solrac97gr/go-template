package application

import (
	"github.com/solrac97gr/go-template/internals/users/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type UserApplication struct {
	repo   domain.UserRepository
	logger *logger.Logger
}

var _ domain.UserApplication = (*UserApplication)(nil)

func NewUserApplication(repo domain.UserRepository, logger *logger.Logger) *UserApplication {
	return &UserApplication{
		repo:   repo,
		logger: logger,
	}
}
