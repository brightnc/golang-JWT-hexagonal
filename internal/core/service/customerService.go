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

func (r *CustomerService) CreateUser(user *domain.IUser) (*domain.IUser, error) {
	customer, err := r.custRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	custResponse := domain.IUser{
		Username:  customer.Username,
		Password:  customer.Password,
		Email:     customer.Email,
		CreatedAt: customer.CreatedAt,
	}
	return &custResponse, nil
}

func (r *CustomerService) UpdateUser(id int, user *domain.IUser) (*domain.UUser, error) {
	customer, err := r.custRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	custResponse := domain.UUser{
		Username: customer.Username,
		Password: customer.Password,
		Email:    customer.Email,
	}
	return &custResponse, nil
}

func (r *CustomerService) DeleteUser(id int) error {
	err := r.custRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerService) FindUser(email string) (*domain.UUser, bool) {
	customer, isExist := r.custRepo.FindUser(email)
	if !isExist {
		return nil, isExist
	}
	return customer, isExist
}
func (r *CustomerService) ListUsers() ([]domain.IUser, error) {
	customersRepo, err := r.custRepo.ListUsers()

	if err != nil {
		return nil, err
	}

	custResponses := []domain.IUser{}
	for _, customer := range customersRepo {
		custResponse := domain.IUser{
			Username:  customer.Username,
			Email:     customer.Email,
			Password:  customer.Password,
			CreatedAt: customer.CreatedAt,
		}
		custResponses = append(custResponses, custResponse)
	}
	return custResponses, nil
}
