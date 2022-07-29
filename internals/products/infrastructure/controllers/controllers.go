package controllers

import (
	"github.com/solrac97gr/go-template/internals/products/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type ProductControllers struct {
	app    domain.ProductApplication
	logger *logger.Logger
}

func NewProductControllers(app domain.ProductApplication, logger *logger.Logger) *ProductControllers {
	return &ProductControllers{
		app:    app,
		logger: logger,
	}
}
