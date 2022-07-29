package controllers

import (
	"github.com/solrac97gr/go-template/internals/payments/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type PaymentControllers struct {
	app    domain.PaymentApplication
	logger *logger.Logger
}

func NewPaymentControllers(app domain.PaymentApplication, logger *logger.Logger) *PaymentControllers {
	return &PaymentControllers{
		app:    app,
		logger: logger,
	}
}
