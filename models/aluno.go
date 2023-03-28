package models

import (

)

type Aluno struct {
	UniqueID  				uint   					`json:"unique_id" gorm:"primaryKey"`
	Nome      				string  				`json:"nome" gorm:"not null" validate:"required"`
	Cpf      				string  				`json:"cpf" gorm:"not null;unique;column:cpf" validate:"required"`
	DisciplinaMatriculas 	[]DisciplinaMatricula	`json:"disciplina_matriculas"`
}