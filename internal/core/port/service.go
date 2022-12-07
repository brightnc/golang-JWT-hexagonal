package port

import "playground/internal/core/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.CustomerResponse, error)
	GetCustomerByID(id int) (*domain.CustomerResponse, error)
	CreateCustomer(c *domain.Customer) (*domain.Customer, error)
}
