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
	var response Contracts.GetAllEjerciciosResponse

	ejercicios, err := ec.service.GetAllEjercicios()
	if err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al obtener ejercicios: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response = Contracts.GetAllEjerciciosResponse{
		BaseResponse: Contracts.BaseResponse{Success: true, Message: "Ejercicios obtenidos correctamente"},
		Ejercicios:   ejercicios,
	}

	c.JSON(http.StatusOK, response)
}

func (ec *EjercicioController) GetEjercicioByID(c *gin.Context) {
	var request Contracts.GetEjercicioByIdRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al obtener el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	ejercicio, err := ec.service.GetEjercicioByID(request.ID)
	if err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al obtener el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, ejercicio)
}

func (ec *EjercicioController) CreateEjercicio(c *gin.Context) {
	var request Contracts.CreateEjercicioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al crear el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	ejercicio := Models.Ejercicio{
		Descripcion: request.Descripcion,
	}

	if err := ec.service.CreateEjercicio(ejercicio); err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al crear el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := Contracts.BaseResponse{Success: true, Message: "Ejercicio creado correctamente"}
	c.JSON(http.StatusInternalServerError, response)
}

func (ec *EjercicioController) DeleteEjercicio(c *gin.Context) {
	var request Contracts.DeleteEjercicioRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al borrar el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	err := ec.service.DeleteEjercicio(request.ID)
	if err != nil {
		response := Contracts.BaseResponse{Success: false, Message: "Error al borrar el ejercicio: " + err.Error()}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response := Contracts.BaseResponse{Success: true, Message: "Ejercicio eliminado correctamente"}
	c.JSON(http.StatusInternalServerError, response)
}
