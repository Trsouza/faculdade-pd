package models

import (
	"time"
)

type DisciplinaMatricula struct {
	UniqueID  					uint   				 `json:"unique_id" gorm:"primaryKey"`
	DataMatricula 				time.Time			 `json:"data_matricula"`
	CursoDisciplinaUniqueID 	uint   				 `json:"curso_disciplina_unique_id" gorm:"not null"`
	AlunoUniqueID 				uint    	 		 `json:"aluno_unique_id" gorm:"not null"`
	CursoDisciplina      		CursoDisciplina  	 `json:"curso_disciplina" gorm:"foreignkey:CursoDisciplinaUniqueID"`
	Aluno      					Aluno  		 		 `json:"aluno" gorm:"foreignkey:AlunoUniqueID"`

}