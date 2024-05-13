package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"rutina.api/Contracts"
	"rutina.api/Models"
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

	router.GET("/ejercicios/get", controller.GetAllEjercicios)
	router.GET("/ejercicios/getById", controller.GetEjercicioByID)
	router.POST("/ejercicios/create", controller.CreateEjercicio)
	router.DELETE("/ejercicios/delete", controller.DeleteEjercicio)
}

func (ec *EjercicioController) GetAllEjercicios(c *gin.Context) {
	ejercicios, err := ec.service.GetAllEjercicios()
	if err != nil {
		c.Error(err)
		return
	}

	response := Contracts.GetAllEjerciciosResponse{
		BaseResponse: Contracts.BaseResponse{Success: true, Message: "Ejercicios obtenidos correctamente"},
		Ejercicios:   ejercicios,
	}

	c.JSON(http.StatusOK, response)
}

func (ec *EjercicioController) GetEjercicioByID(c *gin.Context) {
	var request Contracts.GetEjercicioByIdRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.Error(err)
		return
	}

	ejercicio, err := ec.service.GetEjercicioByID(request.ID)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, ejercicio)
}

func (ec *EjercicioController) CreateEjercicio(c *gin.Context) {
	var request Contracts.CreateEjercicioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	ejercicio := Models.Ejercicio{
		Descripcion: request.Descripcion,
	}

	if err := ec.service.CreateEjercicio(ejercicio); err != nil {
		c.Error(err)
		return
	}

	response := Contracts.BaseResponse{Success: true, Message: "Ejercicio creado correctamente"}
	c.JSON(http.StatusCreated, response)
}

func (ec *EjercicioController) DeleteEjercicio(c *gin.Context) {
	var request Contracts.DeleteEjercicioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := ec.service.DeleteEjercicio(request.ID); err != nil {
		c.Error(err)
		return
	}

	response := Contracts.BaseResponse{Success: true, Message: "Ejercicio eliminado correctamente"}
	c.JSON(http.StatusOK, response)
}
