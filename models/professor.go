package models

import (
)

type Professor struct {
	UniqueID  	 	uint   			`json:"unique_id" gorm:"primaryKey"`
	Nome			string  		`json:"nome" gorm:"not null" validate:"required"`
	Formacao		string  		`json:"formacao" gorm:"not null" validate:"required"`
	Disciplinas		[]Disciplina 	`json:"disciplinas"`
}