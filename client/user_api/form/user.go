package form

type User struct {
	Username string `validate:"required"  binding:"required"`
	Password string `validate:"required"  binding:"required"`
	Email    string `validate:"required"  binding:"required"`
	Phone    string `validate:"required"  binding:"required"`
}

type Account struct {
	UserID   int64   `validate:"required"  binding:"required"`
	Currency string  `validate:"required"  binding:"required"`
	Balance  float64 `validate:"required"  binding:"required"`
}

type Transfer struct {
	FromID        int64   `validate:"required"  binding:"required"`
	ToID          int64   `validate:"required"  binding:"required"`
	Amount        float64 `validate:"required"  binding:"required"`
	AccountNumber string  `validate:"required"  binding:"required"`
}
