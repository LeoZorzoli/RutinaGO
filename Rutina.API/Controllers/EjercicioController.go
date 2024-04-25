package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rutina.api/Repositories"
	"rutina.api/Services"
)

type EjercicioController struct {
	service *Services.EjercicioService
}

func NewEjercicioController(repository Repositories.EjercicioRepository) *EjercicioController {
	service := Services.NewEjercicioService(repository)
	return &EjercicioController{service: service}
}

func RegisterEjercicioRoutes(router *gin.Engine, repository Repositories.EjercicioRepository) {
	controller := NewEjercicioController(repository)
	router.GET("/ejercicios", controller.handleEjercicioRequest)
}

func (ec *EjercicioController) handleEjercicioRequest(c *gin.Context) {
	ejercicios, err := ec.service.GetAllEjercicios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener ejercicios"})
		return
	}
	c.JSON(http.StatusOK, ejercicios)
}
