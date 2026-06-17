package dto

import "time"

type RegisterReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"min=6"`
}

type RegisterRes struct {
	Id        int        `json:"id"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
}
