package Contracts

import "rutina.api/Models"

type GetAllEjerciciosResponse struct {
	BaseResponse
	Ejercicios []Models.Ejercicio `json:"ejercicios"`
}
