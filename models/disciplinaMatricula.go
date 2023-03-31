package models

import (
	"time"
)

type DisciplinaMatricula struct {
	UniqueID  					uint   				 `json:"unique_id" gorm:"primaryKey"`
	DataMatricula 				time.Time			 `json:"data_matricula" gorm:"not null"`
	CursoDisciplinaUniqueID 	uint   				 `json:"curso_disciplina_unique_id" gorm:"not null;uniqueIndex:idx_curso_disciplina_aluno"`
	AlunoUniqueID 				uint    	 		 `json:"aluno_unique_id" gorm:"not null;uniqueIndex:idx_curso_disciplina_aluno"`
	CursoDisciplina      		CursoDisciplina  	 `json:"curso_disciplina" gorm:"foreignkey:CursoDisciplinaUniqueID"`
	Aluno      					Aluno  		 		 `json:"aluno" gorm:"foreignkey:AlunoUniqueID"`

}