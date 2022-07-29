package application

import (
	"github.com/solrac97gr/go-template/internals/products/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type ProductApplication struct {
	repo   domain.ProductRepository
	logger *logger.Logger
}

var _ domain.ProductApplication = (*ProductApplication)(nil)

func NewProductApplication(repo domain.ProductRepository, logger *logger.Logger) *ProductApplication {
	return &ProductApplication{
		repo:   repo,
		logger: logger,
	}
}
