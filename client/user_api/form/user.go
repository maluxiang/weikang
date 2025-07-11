package form

type User struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Phone    string `json:"phone"  binding:"required"`
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
