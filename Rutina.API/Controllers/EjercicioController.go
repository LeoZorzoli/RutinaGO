package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rutina.api/Services"
)

type EjercicioController struct {
	service Services.EjercicioService
}

func Controller(service Services.EjercicioService) *EjercicioController {
	return &EjercicioController{service: service}
}

func RegisterEjercicioRoutes(router *gin.Engine) {
	router.GET("/ejercicios", handleEjercicioRequest)
}

func handleEjercicioRequest(c *gin.Context) {
	service := &Services.EjercicioService{}

	ejercicios, err := service.GetAllEjercicios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ejercicios"})
		return
	}

	c.JSON(http.StatusOK, ejercicios)
}
