package Repositories

import (
	"database/sql"
	"log"

	"rutina.api/Database"
	"rutina.api/Models"
)

// EjercicioRepository maneja las operaciones relacionadas con los ejercicios en la base de datos.
type EjercicioRepository struct {
	db *sql.DB
}

// NewEjercicioRepository crea una nueva instancia del repositorio de ejercicios.
func NewEjercicioRepository() (*EjercicioRepository, error) {
	log.Println("Iniciando NewEjercicioRepository")
	db, err := Database.InitDB()
	if err != nil {
		log.Println("Error al inicializar la base de datos:", err)
		return nil, err
	}

	log.Println("NewEjercicioRepository inicializado correctamente")

	return &EjercicioRepository{db: db}, nil
}

// GetEjercicios devuelve una lista de ejercicios desde la base de datos.
func (repository *EjercicioRepository) GetAllEjercicios() ([]Models.Ejercicio, error) {

	log.Println("Iniciando GetAllEjercicios")

	query := "SELECT * FROM ejercicio"

	// Ejecutar la consulta
	rows, err := repository.db.Query(query)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		return nil, err
	}
	defer rows.Close()

	// Crear una lista para almacenar los ejercicios escaneados
	var ejercicios []Models.Ejercicio

	// Recorrer las filas y escanear cada una en un struct Ejercicio
	for rows.Next() {
		var ejercicio Models.Ejercicio
		if err := rows.Scan(&ejercicio.Id, &ejercicio.Descripcion, &ejercicio.FechaAlta, &ejercicio.FechaBaja); err != nil {
			log.Println("Error al escanear fila:", err)
			return nil, err
		}
		// Agregar el ejercicio escaneado a la lista
		ejercicios = append(ejercicios, ejercicio)
	}

	// Verificar errores de rows.Next()
	if err := rows.Err(); err != nil {
		log.Println("Error al iterar sobre filas:", err)
		return nil, err
	}

	return ejercicios, nil
}
