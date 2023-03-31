package models

import (

)

type CursoDisciplina struct {
	UniqueID  					uint   					 `json:"unique_id" gorm:"primaryKey"`
	CursoUniqueID 				uint   		 			 `json:"curso_unique_id" gorm:"not null;uniqueIndex:idx_curso_disciplina"`
	DisciplinaUniqueID 			uint    				 `json:"disciplina_unique_id" gorm:"not null;uniqueIndex:idx_curso_disciplina"`
	Curso      					Curso  					 `json:"curso" gorm:"foreignkey:CursoUniqueID"`
	Disciplina      			Disciplina  	 		 `json:"disciplina" gorm:"foreignkey:DisciplinaUniqueID"`
	DisciplinaMatriculas		[]DisciplinaMatricula	 `json:"disciplina_matriculas"`

}