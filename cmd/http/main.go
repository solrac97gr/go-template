package main

import (
	payApplication "github.com/solrac97gr/go-template/internals/payments/application"
	payControllers "github.com/solrac97gr/go-template/internals/payments/infrastructure/controllers"
	payRepositories "github.com/solrac97gr/go-template/internals/payments/infrastructure/repositories"
	prodApplication "github.com/solrac97gr/go-template/internals/products/application"
	prodControllers "github.com/solrac97gr/go-template/internals/products/infrastructure/controllers"
	prodRepositories "github.com/solrac97gr/go-template/internals/products/infrastructure/repositories"
	userApplication "github.com/solrac97gr/go-template/internals/users/application"
	userControllers "github.com/solrac97gr/go-template/internals/users/infrastructure/controllers"
	userRepositories "github.com/solrac97gr/go-template/internals/users/infrastructure/repositories"
	"github.com/solrac97gr/go-template/pkg/config"
	dh "github.com/solrac97gr/go-template/pkg/database_helper"
	"github.com/solrac97gr/go-template/pkg/logger"
	"github.com/solrac97gr/go-template/pkg/server"
	"github.com/solrac97gr/go-template/pkg/validator"
)

func main() {
	// Initialize Packages
	validator := validator.NewValidator()
	config := config.NewConfig(validator)
	logger := logger.NewLogger()

	// Get One DB connection
	db := dh.CreateUniqueDatabaseConnection(config, logger)

	// Create Repositories
	paymentRepository := payRepositories.NewPostgresPaymentsRepository(db, config, logger)
	productRepository := prodRepositories.NewPostgresProductRepository(db, config, logger)
	userRepository := userRepositories.NewPostgresUserRepository(db, config, logger)

	// Create Applications
	paymentApplication := payApplication.NewPaymentApplication(paymentRepository, logger)
	productApplication := prodApplication.NewProductApplication(productRepository, logger)
	userApplication := userApplication.NewUserApplication(userRepository, logger)

	// Create Controllers
	paymentControllers := payControllers.NewPaymentControllers(paymentApplication, logger)
	productControllers := prodControllers.NewProductControllers(productApplication, logger)
	userControllers := userControllers.NewUserControllers(userApplication, logger)

	server.NewServer(
		logger,
		config,
		validator,
		paymentControllers,
		productControllers,
		userControllers,
	)
}
