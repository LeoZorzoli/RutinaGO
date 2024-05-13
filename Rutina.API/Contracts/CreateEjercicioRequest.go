package Contracts

type CreateEjercicioRequest struct {
	Descripcion string `json:"descripcion" validate:"required"`
}
