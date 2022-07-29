package main

import (
	payApplication "github.com/solrac97gr/go-template/internals/payments/application"
	payControllers "github.com/solrac97gr/go-template/internals/payments/infrastructure/controllers"
	payRepositories "github.com/solrac97gr/go-template/internals/payments/infrastructure/repositories"
	prodApplication "github.com/solrac97gr/go-template/internals/products/application"
	prodControllers "github.com/solrac97gr/go-template/internals/products/infrastructure/controllers"
	prodRepositories "github.com/solrac97gr/go-template/internals/products/infrastructure/repositories"
	dh "github.com/solrac97gr/go-template/pkg/DatabaseHelper"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
	"github.com/solrac97gr/go-template/pkg/server"
	"github.com/solrac97gr/go-template/pkg/validator"
)

func main() {
	validator := validator.NewValidator()
	config := config.NewConfig(validator)
	logger := logger.NewLogger()

	db := dh.CreateUniqueDatabaseConnection(config, logger)

	paymentRepository := payRepositories.NewPostgresPaymentsRepository(db, config, logger)
	productRepository := prodRepositories.NewPostgresProductRepository(db, config, logger)

	paymentApplication := payApplication.NewPaymentApplication(paymentRepository, logger)
	productApplication := prodApplication.NewProductApplication(productRepository, logger)

	paymentControllers := payControllers.NewPaymentControllers(paymentApplication, logger)
	productControllers := prodControllers.NewProductControllers(productApplication, logger)

	server.NewServer(
		logger,
		config,
		validator,
		paymentControllers,
		productControllers,
	)
}
