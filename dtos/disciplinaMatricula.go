package dtos

import (

)

type DisciplinaMatriculaCreateDTO struct {
	AlunoUniqueID				uint	`json:"aluno_unique_id" validate:"required"`
	CursoDisciplinaUniqueID		uint	`json:"curso_disciplina_unique_i" validate:"required"`
}

type DisciplinaMatriculaUpdateDTO struct {
	AlunoUniqueID				uint	`json:"aluno_unique_id,omitempty"`
	CursoDisciplinaUniqueID 	uint	`json:"curso_disciplina_unique_i,omitempty"`
}