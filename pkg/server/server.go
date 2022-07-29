package server

import (
	"fmt"
	payControllers "github.com/solrac97gr/go-template/internals/payments/infrastructure/controllers"
	prodControllers "github.com/solrac97gr/go-template/internals/products/infrastructure/controllers"
	"github.com/solrac97gr/go-template/pkg/config"
	"github.com/solrac97gr/go-template/pkg/logger"
	"github.com/solrac97gr/go-template/pkg/validator"
)

type Server struct {
	logger             *logger.Logger
	config             *config.Config
	validator          *validator.Validator
	paymentControllers *payControllers.PaymentControllers
	productControllers *prodControllers.ProductControllers
}

func NewServer(
	logger *logger.Logger,
	config *config.Config,
	validator *validator.Validator,
	paymentControllers *payControllers.PaymentControllers,
	productControllers *prodControllers.ProductControllers,

) *Server {
	return &Server{
		logger:             logger,
		config:             config,
		validator:          validator,
		paymentControllers: paymentControllers,
		productControllers: productControllers,
	}
}

func (s *Server) Init() {
	fmt.Println("Starting server")
}
