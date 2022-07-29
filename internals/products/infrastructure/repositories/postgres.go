package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-template/internals/products/domain"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type PostgresProducts struct {
	db     *sql.DB
	config *config.Config
	logger *logger.Logger
}

var _ domain.ProductRepository = (*PostgresProducts)(nil)

func NewPostgresProductRepository(db *sql.DB, config *config.Config, logger *logger.Logger) *PostgresProducts {
	return &PostgresProducts{
		db:     db,
		config: config,
		logger: logger,
	}
}
