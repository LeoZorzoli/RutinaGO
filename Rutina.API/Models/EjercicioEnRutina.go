package Models

import "time"

type EjercicioEnRutina struct {
	Ejercicio        Ejercicio `json:"ejercicio"`
	DiaEntrenamiento time.Time `json:"diaEntrenamiento"`
	Series           int       `json:"series"`
	Pesos            int       `json:"pesos"`
}
