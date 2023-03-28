package models

import (

)

type Faculdade struct {
	UniqueID  uint    `json:"unique_id" gorm:"primaryKey"`
	Nome      string  `json:"nome" gorm:"not null" validate:"required"`
	Cnpj      string  `json:"cnpj" gorm:"not null;unique;column:cnpj" validate:"required"`
	Cursos 	  []Curso `json:"cursos"`
}