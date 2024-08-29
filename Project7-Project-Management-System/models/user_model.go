package models

type UserModel struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
