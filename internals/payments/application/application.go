package application

import (
	"github.com/solrac97gr/go-template/internals/payments/domain"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type PaymentApplication struct {
	repo   domain.PaymentRepository
	logger *logger.Logger
}

var _ domain.PaymentApplication = (*PaymentApplication)(nil)

func NewPaymentApplication(repo domain.PaymentRepository, logger *logger.Logger) *PaymentApplication {
	return &PaymentApplication{
		repo:   repo,
		logger: logger,
	}
}
