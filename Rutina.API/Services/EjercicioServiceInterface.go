package Services

import (
	"rutina.api/Models"
)

type EjercicioServiceInterface interface {
	GetAllEjercicios() ([]Models.Ejercicio, error)
	GetEjercicioByID(id int) (Models.Ejercicio, error)
	CreateEjercicio(ejercicio Models.Ejercicio) error
	UpdateEjercicio(id int, ejercicio Models.Ejercicio) error
	DeleteEjercicio(id int) error
}
