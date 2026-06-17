package dto

type ErrorResponse struct {
	Message string `json:"message"`
	Success bool   `json:"isSucces"`
	Error   string `json:"error"`
}

type Response[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
	Success bool   `json:"isSucces"`
}
