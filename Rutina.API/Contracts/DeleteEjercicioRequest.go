package Contracts

type DeleteEjercicioRequest struct {
	ID int `json:"ID" validate:"required"`
}
