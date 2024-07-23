package form

type Points struct {
	UserID int64 `binding:"required"`
	Points int64 `binding:"required"`
}
