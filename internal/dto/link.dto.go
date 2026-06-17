package dto

import "time"

type CreateLinkReq struct {
	OriginalUrl string `json:"original_url" binding:"required,url"`
	Slug        string `json:"slug" binding:"omitempty"`
}

type LinkDetailRes struct {
	Id          int       `json:"id"`
	OriginalUrl string    `json:"original_url"`
	Slug        string    `json:"slug"`
	ShortUrl    string    `json:"short_url"`
	ClickCount  int       `json:"click_count"`
	CreatedAt   time.Time `json:"created_at"`
}
