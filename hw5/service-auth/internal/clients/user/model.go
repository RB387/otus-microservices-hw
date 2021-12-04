package user

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CheckPassword struct {
	Result Result `json:"result"`
}

type Result struct {
	Valid bool `json:"valid"`
}
