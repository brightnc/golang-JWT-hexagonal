package repository

import (
	"fmt"
	"playground/internal/core/domain"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type customerRepositoryDB struct {
	db *sqlx.DB
}

func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{db: db}
}

// func (r *customerRepositoryDB) GetAll() ([]domain.Customer, error) {
// 	customers := []domain.Customer{}
// 	query := "select customer_id, user_name, password, email from customers"
// 	err := r.db.Select(&customers, query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return customers, nil
// }

// func (r *customerRepositoryDB) GetByID(id int) (*domain.Customer, error) {
// 	customer := domain.Customer{}
// 	query := "select customer_id, user_name, password, email from customers where customer_id=?"
// 	err := r.db.Get(&customer, query, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &customer, nil
// }

func (r *customerRepositoryDB) CreateUser(user *domain.IUser) (*domain.IUser, error) {
	query := "insert into customers (user_name, password, email, createdAt) values (?, ?,?, ?)"
	hasedByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	currentTime := time.Now()
	dateNow := fmt.Sprintf("%d-%d-%d %d:%d:%d", currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Hour(), currentTime.Second())
	user.CreatedAt = dateNow
	user.Password = string(hasedByte)
	_, err = r.db.Exec(query, user.Username, user.Password, user.Email, user.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}

func (r *customerRepositoryDB) UpdateUser(id int, user *domain.IUser) (*domain.IUser, error) {
	query := "update customers set user_name = ?, password = ?, email = ? where customer_id=?"
	hasedByte, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	user.Password = string(hasedByte)
	_, err = r.db.Exec(query, user.Username, user.Password, user.Email, id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func (r *customerRepositoryDB) DeleteUser(id int) error {
	query := "delete from customers where customer_id = ?"

	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *customerRepositoryDB) FindUser(email string) (*domain.UUser, bool) {
	customer := domain.UUser{}
	query := "select customer_id, user_name, password, email from customers where email = ?"
	err := r.db.Get(&customer, query, email)
	if err != nil {
		return nil, false
	}
	return &customer, true

}

func (r *customerRepositoryDB) ListUsers() ([]domain.IUser, error) {
	customers := []domain.IUser{}
	query := "select user_name, password, email, createdAt from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
