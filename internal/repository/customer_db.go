package repository

import (
	"playground/internal/core/domain"

	"github.com/jmoiron/sqlx"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

func (r *customerRepositoryDB) GetAll() ([]domain.Customer, error) {
	customers := []domain.Customer{}
	query := "select customer_id, user_name, password, email from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r *customerRepositoryDB) GetByID(id int) (*domain.Customer, error) {
	customer := domain.Customer{}
	query := "select customer_id, user_name, password, email from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (r *customerRepositoryDB) CreateCustomer(c *domain.Customer) (*domain.Customer, error) {
	query := "insert into customers (user_name, password, email) values (?, ?,?)"
	_, err := r.db.Exec(query, c.Username, c.Password, c.Email)
	if err != nil {
		return nil, err
	}
	return c, nil
}
