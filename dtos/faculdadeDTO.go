package dtos

import (

)

type FaculdadeCreateDTO struct {
	Nome	string  `json:"nome" validate:"required"`
	Cnpj	string  `json:"cnpj" validate:"required"`
}

type FaculdadeUpdateDTO struct {
	Nome	string  `json:"nome,omitempty"`
	Cnpj	string  `json:"cnpj,omitempty"`
}