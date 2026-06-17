package model

import "time"

type User struct {
	Id        int        `db:"id"`
	Email     string     `db:"email"`
	Password  string     `db:"password"`
	Name      *string    `db:"name"`
	Job       *string    `db:"job"`
	Photo     *string    `db:"photo"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
