package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./Database/rutina.db"

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	log.Println("Conexión a la base de datos establecida")

	return db, nil
}
