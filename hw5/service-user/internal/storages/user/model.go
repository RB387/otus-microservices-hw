package user

type User struct {
	Username  string `db:"username"`
	Password  string `db:"password"`
	Salt      string `db:"salt"`
	FirstName string `db:"firstname"`
	LastName  string `db:"lastname"`
	Email     string `db:"email"`
	Phone     string `db:"phone"`
}
