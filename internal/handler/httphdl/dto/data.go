package dto

import "playground/internal/core/domain"

type Customer struct {
	Username string `json:"user_name"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email"`
}

func (r *Customer) ToDomainCustomer() domain.Customer {
	return domain.Customer{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}
}
