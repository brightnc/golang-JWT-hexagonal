package port

import "playground/internal/core/domain"

type CustomerRepository interface {
	GetAll() ([]domain.Customer, error)
	GetByID(id int) (*domain.Customer, error)
	CreateCustomer(c *domain.Customer) (*domain.Customer, error)
}
