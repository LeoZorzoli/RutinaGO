package models
import "time"

type Ejercicio struct {
	Id 			int 		`json:"id"`
	Descripcion string  	`json:"descripcion"`
	FechaAlta 	time.Time 	`json:"fechaAlta"`
	FechaBaja 	time.Time 	`json:"fechaBaja"`
}