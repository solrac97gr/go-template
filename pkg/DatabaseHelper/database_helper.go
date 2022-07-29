package database_helper

import (
	"database/sql"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
)

var (
	database *sql.DB
)

func CreateUniqueDatabaseConnection(config *config.Config, logger *logger.Logger) *sql.DB {
	if database != nil {
		env := config.GetEnvironmentVariables()
		db, err := sql.Open("postgres", env.DbUrl)
		if err != nil {
			logger.Error(err)
			panic(err)
		}
		database = db
	}
	return database
}
