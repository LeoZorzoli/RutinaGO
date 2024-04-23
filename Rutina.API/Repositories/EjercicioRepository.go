package Repositories

import (
	"database/sql"

	"rutina.api/Database"
	"rutina.api/Models"
)

// EjercicioRepository maneja las operaciones relacionadas con los ejercicios en la base de datos.
type EjercicioRepository struct {
	db *sql.DB
}

// NewEjercicioRepository crea una nueva instancia del repositorio de ejercicios.
func NewEjercicioRepository() *EjercicioRepository {
	db, err := Database.InitDB()
	if err != nil {
		panic(err)
	}

	return &EjercicioRepository{db: db}
}

// GetEjercicios devuelve una lista de ejercicios desde la base de datos.
func (repository *EjercicioRepository) GetAllEjercicios() ([]Models.Ejercicio, error) {

	query := "SELECT * FROM ejercicio"
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var ejercicios []Models.Ejercicio
	for rows.Next() {
		var ejercicio Models.Ejercicio
		err := rows.Scan(&ejercicio.Id, &ejercicio.Descripcion)
		if err != nil {
			return nil, err
		}
		ejercicios = append(ejercicios, ejercicio)
	}

	return ejercicios, nil
}
