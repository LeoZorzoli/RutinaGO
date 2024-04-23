package Database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "./rutina.db"

func InitDB() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
		return nil, err
	}

	log.Println("Conexión a la base de datos establecida")

	defer db.Close()
	return db, nil
}
