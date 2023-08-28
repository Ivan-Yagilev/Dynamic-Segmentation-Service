package segmentation_service

type User struct {
	Id       int    `json:"-" db:"id"`
	Username string `json:"username" db:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
