package main

import (
	"github.com/gin-gonic/gin"
	"rutina.api/Controllers"
)

func main() {
	router := gin.Default()
	Controllers.RegisterEjercicioRoutes(router)

	err := router.Run(":8080")
	if err != nil {
		panic("Error al iniciar el servidor: " + err.Error())
	}
}
