package domain

type PaymentApplication interface {
	Create(p *Payment) error
	Delete(id string) error
	Update(id string) error
	GetById(id string) *Payment
	GetAll() []*Payment
}
