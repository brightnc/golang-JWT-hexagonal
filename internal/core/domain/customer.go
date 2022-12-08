package domain

type IUser struct {
	Username  string `db:"user_name"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	CreatedAt string `db:"createdAt"`
}

type UUser struct {
	Id       string `db:"customer_id"`
	Username string `db:"user_name"`
	Password string `db:"password"`
	Email    string `db:"email"`
}
