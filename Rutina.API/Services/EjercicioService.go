package Services

import (
	"rutina.api/Models"
	"rutina.api/Repositories"
)

type EjercicioService struct {
	repository Repositories.EjercicioRepository
}

var _ EjercicioServiceInterface = (*EjercicioService)(nil)

func NewEjercicioService(repository Repositories.EjercicioRepository) *EjercicioService {
	return &EjercicioService{repository: repository}
}

func (service *EjercicioService) GetAllEjercicios() ([]Models.Ejercicio, error) {
	return service.repository.GetAllEjercicios()
}

func (service *EjercicioService) GetEjercicioByID(id int) (Models.Ejercicio, error) {

	return Models.Ejercicio{}, nil
}

func (service *EjercicioService) CreateEjercicio(ejercicio Models.Ejercicio) error {
	return nil
}

func (service *EjercicioService) UpdateEjercicio(id int, ejercicio Models.Ejercicio) error {
	return nil
}

func (service *EjercicioService) DeleteEjercicio(id int) error {
	return nil
}
