package models

import (

)

type Disciplina struct {
	UniqueID  			uint    			`json:"unique_id" gorm:"primaryKey"`
	Nome      			string  			`json:"nome" gorm:"not null"`
	ProfessorUniqueID 	uint    			`json:"professor_unique_id" gorm:"not null"`
	Professor      		Professor   		`json:"professor" gorm:"foreignkey:ProfessorUniqueID"`
	CursoDisciplinas	[]CursoDisciplina   `json:"curso_disciplinas"`
}