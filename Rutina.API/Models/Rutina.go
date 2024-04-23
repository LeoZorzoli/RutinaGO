package Models

import "time"

type Rutina struct {
	Id          int                 `json:"id"`
	Descripcion string              `json:"descripcion"`
	FechaInicio time.Time           `json:"fechaInicio"`
	FechaFin    time.Time           `json:"fechaFin"`
	FechaAlta   time.Time           `json:"fechaAlta"`
	FechaBaja   time.Time           `json:"fechaBaja"`
	Ejercicios  []EjercicioEnRutina `json:"ejercicios"`
}
