package usecase

type SistemaExpertoUsecase interface {
	//firmas de los metodos
}

type sistemaExpertoUseCase struct {
	//atributos (si los hay) por ejemplo: // db *sql.DB // logger *log.Logger
}

func NewSistemaExpertoUseCase() SistemaExpertoUsecase {
	return &sistemaExpertoUseCase{}
}

//funciones de los metodos
