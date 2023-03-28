package migrations

import (
	"github.com/trsouza/faculdade-pd/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Curso{})
	db.AutoMigrate(models.Faculdade{})
	db.AutoMigrate(models.Disciplina{})
	db.AutoMigrate(models.Professor{})
	db.AutoMigrate(models.Aluno{})
	db.AutoMigrate(models.DisciplinaMatricula{})
	db.AutoMigrate(models.CursoDisciplina{})
}