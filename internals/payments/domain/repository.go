package domain

type PaymentRepository interface {
	Create(p *Payment) error
	Delete(id string) error
	Update(id string) error
	GetById(id string) *Payment
	GetAll() []*Payment
}
