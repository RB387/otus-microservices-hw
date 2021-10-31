package user

type User struct {
	Username  string `db:"username"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}
