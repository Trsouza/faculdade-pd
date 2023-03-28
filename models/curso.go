package models

import (

)

type Curso struct {
	UniqueID  				uint   				`json:"unique_id" gorm:"primaryKey"`
	Nome      				string  			`json:"nome" gorm:"not null"`
	FaculdadeUniqueID 		uint    			`json:"faculdade_unique_id" gorm:"not null"`
	Faculdade      			Faculdade   		`json:"faculdade" gorm:"foreignkey:FaculdadeUniqueID"`
	CursoDisciplinas		[]CursoDisciplina   `json:"curso_disciplinas"`
}