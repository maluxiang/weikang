package form

type User struct {
	Username string `validate:"required"  binding:"required"`
	Password string `validate:"required"  binding:"required"`
	Email    string `validate:"required"  binding:"required"`
	Phone    string `validate:"required"  binding:"required"`
}
