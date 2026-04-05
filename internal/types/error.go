package types

type AppError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}
