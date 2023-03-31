package dtos

import (

)

type AlunoCreateDTO struct {
	Nome	string  `json:"nome" validate:"required"`
	Cpf		string  `json:"cpf" validate:"required"`
}

type AlunoUpdateDTO struct {
	Nome	string  `json:"nome,omitempty"`
	Cpf		string  `json:"cpf,omitempty"`
}