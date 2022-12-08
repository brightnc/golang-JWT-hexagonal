package port

import "playground/internal/core/domain"

type CustomerRepository interface {
	CreateUser(user *domain.IUser) (*domain.IUser, error)
	UpdateUser(id int, user *domain.IUser) (*domain.IUser, error)
	DeleteUser(id int) error
	FindUser(email string) (*domain.UUser, bool)
	ListUsers() ([]domain.IUser, error)
}
