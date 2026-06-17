package repository

import (
	"context"
	"log"

	"github.com/iamhanif11/shortlink-backend.git/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository) AddUser(ctx context.Context, email, hashPassword string) (model.User, error) {
	sql := `
		INSERT INTO users (email, password) 
		VALUES ($1, $2)
		RETURNING id, email, created_at
	`
	log.Println(email, hashPassword)
	args := []any{email, hashPassword}

	var user model.User
	if err := ar.db.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Email, &user.CreatedAt); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	log.Println(email)
	sql := `
		SELECT id, email, password, name, photo, job
		FROM users
		WHERE email = $1
	`
	args := []any{email}

	var user model.User
	if err := ar.db.QueryRow(ctx, sql, args...).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.Photo, &user.Job); err != nil {
		return model.User{}, err
	}
	return user, nil
}
