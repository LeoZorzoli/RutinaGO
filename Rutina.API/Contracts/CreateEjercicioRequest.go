package Contracts

import (
	"time"
)

type CreateEjercicioRequest struct {
	Descripcion string     `json:"descripcion" validate:"required"`
	FechaAlta   time.Time  `json:"fechaAlta" validate:"required"`
	FechaBaja   *time.Time `json:"fechaBaja"`
}
