package dto

import "playground/internal/core/domain"

type Customer struct {
	Username  string `json:"user_name"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
}

func (r *Customer) ToDomainCustomer() domain.IUser {
	return domain.IUser{
		Username:  r.Username,
		Password:  r.Password,
		Email:     r.Email,
		CreatedAt: r.CreatedAt,
	}
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
