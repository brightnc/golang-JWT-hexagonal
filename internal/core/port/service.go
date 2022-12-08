package port

import "playground/internal/core/domain"

type CustomerService interface {
	CreateUser(user *domain.IUser) (*domain.IUser, error)
	UpdateUser(id int, user *domain.IUser) (*domain.UUser, error)
	DeleteUser(id int) error
	FindUser(email string) (*domain.UUser, bool)
	ListUsers() ([]domain.IUser, error)
}
