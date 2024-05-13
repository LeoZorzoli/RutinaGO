package Repositories

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"rutina.api/Database"
	"rutina.api/Models"
)

type EjercicioRepository struct {
	db *sql.DB
}

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

func (repository *EjercicioRepository) GetAllEjercicios() ([]Models.Ejercicio, error) {

	log.Println("Iniciando GetAllEjercicios")

	query := "SELECT * FROM ejercicio WHERE eje_fec_baj IS null"

	rows, err := repository.db.Query(query)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		return nil, err
	}
	defer rows.Close()

	var ejercicios []Models.Ejercicio

	for rows.Next() {
		var row Models.EjercicioRow
		if err := rows.Scan(&row.ID, &row.Descripcion, &row.FechaAlta, &row.FechaBaja); err != nil {
			log.Println("Error al escanear fila:", err)
			return nil, err
		}

		var ejercicio Models.Ejercicio
		ejercicio.ID = row.ID
		ejercicio.Descripcion = row.Descripcion
		ejercicio.FechaAlta = row.FechaAlta
		ejercicio.FechaBaja = row.FechaBaja

		ejercicios = append(ejercicios, ejercicio)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error al iterar sobre filas:", err)
		return nil, err
	}

	return ejercicios, nil
}

func (repository *EjercicioRepository) GetEjercicioByID(id int) (Models.Ejercicio, error) {
	log.Println("Iniciando GetEjercicioByID")

	var ejercicio = Models.Ejercicio{}

	query := "SELECT * FROM ejercicio WHERE eje_id = ?"

	var row Models.EjercicioRow
	err := repository.db.QueryRow(query, id).Scan(&row.ID, &row.Descripcion, &row.FechaAlta, &row.FechaBaja)

	ejercicio.ID = row.ID
	ejercicio.Descripcion = row.Descripcion
	ejercicio.FechaAlta = row.FechaAlta
	ejercicio.FechaBaja = row.FechaBaja

	if err != nil {
		if err == sql.ErrNoRows {
			return ejercicio, fmt.Errorf("no se encontró ningún ejercicio con el ID: %d", id)
		}
		log.Println("Error al ejecutar la consulta:", err)
		return ejercicio, err
	}

	return ejercicio, nil
}

func (repository *EjercicioRepository) CreateEjercicio(ejercicio Models.Ejercicio) error {
	log.Println("Iniciando CreateEjercicio")

	row := Models.EjercicioRow{
		Descripcion: ejercicio.Descripcion,
		FechaAlta:   time.Now(),
		FechaBaja:   nil,
	}

	query := "INSERT INTO ejercicio (eje_desc, eje_fec_alt, eje_fec_baj) VALUES (?, ?, ?)"

	_, err := repository.db.Exec(query, row.Descripcion, row.FechaAlta, row.FechaBaja)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		return err
	}

	return nil
}

func (repository *EjercicioRepository) DeleteEjercicio(id int) error {
	log.Println("Iniciando DeleteEjercicio")
	log.Println(id)

	query := "UPDATE ejercicio SET eje_fec_baj = date('now') WHERE eje_id = ?"

	_, err := repository.db.Exec(query, id)
	if err != nil {
		log.Println("Error al ejecutar la consulta:", err)
		return err
	}

	return nil
}
