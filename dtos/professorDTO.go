package dtos

import (

)

type ProfessorCreateDTO struct {
	Nome		string  `json:"nome" validate:"required"`
	Formacao	string  `json:"formacao" validate:"required"`
}

type ProfessorUpdateDTO struct {
	Nome		string  `json:"nome,omitempty"`
	Formacao	string  `json:"formacao,omitempty"`
}