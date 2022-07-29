package repositories

import "github.com/solrac97gr/go-template/internals/payments/domain"

func (repo *PostgresPayments) GetById(id string) *domain.Payment {
	return &domain.Payment{}
}
