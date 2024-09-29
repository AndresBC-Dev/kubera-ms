package usermodel

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;autoIncrement"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
