package dtos

import (

)

type CursoCreateDTO struct {
	Nome				string	`json:"nome" validate:"required"`
	FaculdadeUniqueID	uint	`json:"faculdade_unique_id" validate:"required"`
}

type CursoUpdateDTO struct {
	Nome				string	`json:"nome,omitempty"`
	FaculdadeUniqueID 	uint	`json:"faculdade_unique_id,omitempty"`
}