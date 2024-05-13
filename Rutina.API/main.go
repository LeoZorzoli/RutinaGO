package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"rutina.api/Controllers"
	"rutina.api/Middleware"
	"rutina.api/Repositories"
)

func main() {
	repository, err := InitializeDatabase()
	if err != nil {
		panic("Error al inicializar la base de datos: " + err.Error())
	}

	router := gin.Default()

	router.Use(Middleware.ErrorHandler())

	Controllers.RegisterEjercicioRoutes(router, repository)

	err = router.Run(":8080")
	if err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}

func InitializeDatabase() (Repositories.EjercicioRepository, error) {
	repository, err := Repositories.NewEjercicioRepository()
	if err != nil {
		log.Println("Error al inicializar la base de datos:", err)
	}
	return *repository, err
}
