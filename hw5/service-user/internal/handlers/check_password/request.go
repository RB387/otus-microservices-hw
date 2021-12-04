package check_password

type ValidatePassword struct {
	Password string `json:"password" binding:"required"`
}
