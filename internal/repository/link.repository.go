package repository

import (
	"context"
	"errors"

	"github.com/iamhanif11/shortlink-backend.git/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepository struct {
	db *pgxpool.Pool
}

func NewLinkRepository(db *pgxpool.Pool) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (l *LinkRepository) AddNewLink(ctx context.Context, link model.Link) (model.Link, error) {
	q := `
		INSERT INTO links (user_id, original_url, slug)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, original_url, slug, click_count, created_at
	`

	var newLink model.Link
	err := l.db.QueryRow(ctx, q, link.UserId, link.OriginalUrl, link.Slug).Scan(
		&newLink.Id,
		&newLink.UserId,
		&newLink.OriginalUrl,
		&newLink.Slug,
		&newLink.ClickCount,
		&newLink.CreatedAt,
	)
	if err != nil {
		return model.Link{}, err
	}
	return newLink, nil
}

func (l *LinkRepository) IsSlugExist(ctx context.Context, slug string) (bool, error) {
	q := `
			SELECT 1
			FROM links
			WHERE slug = $1 AND deleted_at IS NULL
			LIMIT 1
		
	`
	var placeholder int
	err := l.db.QueryRow(ctx, q, slug).Scan(&placeholder)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
