package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-template/internals/payments/domain"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type PostgresPayments struct {
	db     *sql.DB
	config *config.Config
	logger *logger.Logger
}

var _ domain.PaymentRepository = (*PostgresPayments)(nil)

func NewPostgresPaymentsRepository(db *sql.DB, config *config.Config, logger *logger.Logger) *PostgresPayments {
	return &PostgresPayments{
		db:     db,
		config: config,
		logger: logger,
	}
}
