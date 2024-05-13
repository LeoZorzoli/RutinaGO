package Models

import "time"

type Ejercicio struct {
	ID          int        `json:"id"`
	Descripcion string     `json:"descripcion"`
	FechaAlta   time.Time  `json:"fechaAlta"`
	FechaBaja   *time.Time `json:"fechaBaja"`
}

type EjercicioRow struct {
	ID          int        `db:"eje_id"`
	Descripcion string     `db:"eje_desc"`
	FechaAlta   time.Time  `db:"eje_fec_alt"`
	FechaBaja   *time.Time `db:"eje_fec_baj"`
}
