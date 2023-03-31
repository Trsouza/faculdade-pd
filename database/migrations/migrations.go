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

	// Seed(db)
}


func Seed(db *gorm.DB) {

	var nome = "Nome = ?";

	alunos := []models.Aluno{
		{ Nome: "Aluno 1", Cpf: "44840756015" },
		{ Nome: "Aluno 2", Cpf: "55198584037" },
	}

	for _, aluno := range alunos {
		db.Create(&aluno)
	}

	var alu1, alu2 models.Aluno
	db.First(&alu1, nome, "Aluno 1")
	db.First(&alu2, nome, "Aluno 2")

	faculdades := []models.Faculdade{
		{ Nome: "Faculdade 1", Cnpj: "72124104000119", },
		{ Nome: "Faculdade 2", Cnpj: "12454370000188", },
	}

	for _, faculdade := range faculdades {
		db.Create(&faculdade)
	}

	var fac1, fac2 models.Faculdade

	db.First(&fac1, nome, "Faculdade 1")
	db.First(&fac2, nome, "Faculdade 2")

	cursos := []models.Curso{
		{ Nome: "Curso 1", FaculdadeUniqueID: fac1.UniqueID },
		{ Nome: "Curso 2", FaculdadeUniqueID: fac2.UniqueID },
		{ Nome: "Curso 3", FaculdadeUniqueID: fac2.UniqueID },
	}

	for _, curso := range cursos {
		db.Create(&curso)
	}

	var curso1, curso2, curso3 models.Curso

	db.First(&curso1, nome, "Curso 1")
	db.First(&curso2, nome, "Curso 2")
	db.First(&curso3, nome, "Curso 3")

	professores := []models.Professor{
		{ Nome: "Professor 1", Formacao: "Doutorado" },
		{ Nome: "Professor 2", Formacao: "Mestrado" },
	}

	for _, professor := range professores {
		db.Create(&professor)
	}

	var prof1, prof2 models.Professor
	db.First(&prof1, nome, "Professor 1")
	db.First(&prof2, nome, "Professor 2")

	disciplinas := []models.Disciplina{
		{ Nome: "Disciplina 1", ProfessorUniqueID: prof1.UniqueID },
		{ Nome: "Disciplina 2", ProfessorUniqueID: prof2.UniqueID },
		{ Nome: "Disciplina 3", ProfessorUniqueID: prof2.UniqueID },
		{ Nome: "Disciplina 4", ProfessorUniqueID: prof2.UniqueID },

	}

	for _, disciplina := range disciplinas {
		db.Create(&disciplina)
	}

}