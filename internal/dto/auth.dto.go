package dto

import "time"

type RegisterReq struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"min=6"`
	ConfirmPassword string `json:"confirm_password" binding://required,eqfield=Password"`
}

type RegisterRes struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required" example:"koda36@gmail.com"`
	Password string `json:"password" binding:"required,min=6" example:"Password123"`
}

type LoginUserDetail struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct {
	Token string          `json:"token"`
	User  LoginUserDetail `json:"user"`
}
