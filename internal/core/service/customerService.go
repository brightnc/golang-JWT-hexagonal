package service

import (
	"playground/internal/core/domain"
	"playground/internal/core/port"
)

type CustomerService struct {
	custRepo port.CustomerRepository
}

func NewCustomerService(custRepo port.CustomerRepository) CustomerService {
	return CustomerService{custRepo: custRepo}
}

func (r *CustomerService) GetAllCustomer() ([]domain.CustomerResponse, error) {
	customersRepo, err := r.custRepo.GetAll()

	if err != nil {
		return nil, err
	}

	custResponses := []domain.CustomerResponse{}
	for _, customer := range customersRepo {
		custResponse := domain.CustomerResponse{
			CustomerID: customer.CustomerID,
			Username:   customer.Email,
			Email:      customer.Email,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil

}

func (r *CustomerService) GetCustomerByID(id int) (*domain.CustomerResponse, error) {
	customerRepo, err := r.custRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	custResponse := domain.CustomerResponse{
		CustomerID: customerRepo.CustomerID,
		Username:   customerRepo.Username,
		Email:      customerRepo.Email,
	}
	return &custResponse, nil

}

func (r *CustomerService) CreateCustomer(c *domain.Customer) (*domain.Customer, error) {
	customer, err := r.custRepo.CreateCustomer(c)
	if err != nil {
		return nil, err
	}
	custResponse := domain.Customer{
		CustomerID: customer.CustomerID,
		Username:   customer.Username,
		Email:      customer.Email,
	}
	return &custResponse, nil
}
