package controller

import (
	usecase "E:/Escritorio/Sistemas-experto/internal/adapters/controllers"
)

type SistemaExpertoController struct {
	sistemaExpertoUseCase usecase.SistemaExpertoUseCase
}

func NewSistemaExpertoController(sistemaExpertoUseCase usecase.SistemaExpertoUseCase) SistemaExpertoController {
	return &sistemaExpertoController{
		sistemaExpertoUseCase: sistemaExpertoUseCase,
	}
}