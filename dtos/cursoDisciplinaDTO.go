package dtos

import (

)

type CursoDisciplinaCreateDTO struct {
	CursoUniqueID 			uint	`json:"curso_unique_id" validate:"required"`
	DisciplinaUniqueID 		uint	`json:"disciplina_unique_id" validate:"required"`
}

type CursoDisciplinaUpdateDTO struct {
	CursoUniqueID 			uint	`json:"curso_unique_id,omitempty"`
	DisciplinaUniqueID 		uint	`json:"disciplina_unique_id,omitempty"`
}