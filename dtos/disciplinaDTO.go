package dtos

import (

)

type DisciplinaCreateDTO struct {
	Nome				string	`json:"nome" validate:"required"`
	ProfessorUniqueID	uint	`json:"professor_unique_id" validate:"required"`
}

type DisciplinaUpdateDTO struct {
	Nome				string	`json:"nome,omitempty"`
	ProfessorUniqueID 	uint	`json:"professor_unique_id,omitempty"`
}