package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-template/internals/users/domain"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
)

type PostgresUsers struct {
	db     *sql.DB
	config *config.Config
	logger *logger.Logger
}

var _ domain.UserRepository = (*PostgresUsers)(nil)

func NewPostgresUserRepository(db *sql.DB, config *config.Config, logger *logger.Logger) *PostgresUsers {
	return &PostgresUsers{
		db:     db,
		config: config,
		logger: logger,
	}
}
