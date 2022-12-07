package domain

type Customer struct {
	CustomerID int    `db:"customer_id"`
	Username   string `db:"user_name"`
	Password   string `db:"password"`
	Email      string `db:"email"`
}
