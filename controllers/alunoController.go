package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
)

// @Summary Filtra alunos por curso
// @Description Realiza um filtro para obter todos os alunos de um curso
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do curso"
// @Success 200 {object} models.Aluno
// @Router /aluno/curso/{id} [GET]
func FiltrarAlunosPorCurso(c *gin.Context) {
	var curso models.Curso
	var alunos []models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&curso, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "curso não encontrado com o id: " + err.Error(),
		})
		return
	}

	db.Table("alunos").
		Joins("JOIN disciplina_matriculas dm ON dm.aluno_unique_id = alunos.unique_id").
		Joins("JOIN curso_disciplinas cd ON cd.unique_id = dm.curso_disciplina_unique_id").
		Joins("JOIN cursos c ON c.unique_id = cd.curso_unique_id").
		Where("c.unique_id = ?", newid).
		Select("DISTINCT alunos.unique_id, alunos.nome, alunos.cpf").
		Find(&alunos)

	c.JSON(200, alunos)
}

// @Summary Filtra alunos por disciplina
// @Description Realiza um filtro para obter todos os alunos de uma disciplina
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do disciplina"
// @Success 200 {object} models.Aluno
// @Router /aluno/disciplina/{id} [GET]
func FiltrarAlunosPorDisciplina(c *gin.Context) {
	var disciplina models.Disciplina
	var alunos []models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&disciplina, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "disciplina não encontrado com o id: " + err.Error(),
		})
		return
	}

	db.Table("alunos").
		Joins("JOIN disciplina_matriculas dm ON dm.aluno_unique_id = alunos.unique_id").
		Joins("JOIN curso_disciplinas cd ON cd.unique_id = dm.curso_disciplina_unique_id").
		Joins("JOIN disciplinas d ON d.unique_id = cd.disciplina_unique_id").
		Where("d.unique_id = ?", newid).
		Select("DISTINCT alunos.unique_id, alunos.nome, alunos.cpf").
		Find(&alunos)

	c.JSON(200, alunos)
}

// @Summary Filtra alunos por faculdade
// @Description Realiza um filtro para obter todos os alunos de uma faculdade
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da faculdade"
// @Success 200 {object} models.Aluno
// @Router /aluno/faculdade/{id} [GET]
func FiltrarAlunosPorFaculdade(c *gin.Context) {
	var faculdade models.Faculdade
	var alunos []models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&faculdade, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "faculdade não encontrado com o id: " + err.Error(),
		})
		return
	}

	db.Table("alunos").
		Joins("JOIN disciplina_matriculas dm ON dm.aluno_unique_id = alunos.unique_id").
		Joins("JOIN curso_disciplinas cd ON cd.unique_id = dm.curso_disciplina_unique_id").
		Joins("JOIN cursos c ON c.unique_id = cd.curso_unique_id").
		Joins("JOIN faculdades f ON f.unique_id = c.faculdade_unique_id").
		Where("f.unique_id = ?", newid).
		Select("DISTINCT alunos.unique_id, alunos.nome, alunos.cpf").
		Find(&alunos)

	c.JSON(200, alunos)
}

// @Summary Filtra o total de alunos de uma faculdade
// @Description Realiza um filtro para obter o total de alunos de uma faculdade
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do faculdade"
// @Success 200 {object} models.Aluno
// @Router /aluno/total/{id} [GET]
func FiltrarTotalAlunosPorFaculdade(c *gin.Context) {
	var faculdade models.Faculdade
	var totalAlunos int64
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&faculdade, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "faculdade não encontrado com o id: " + err.Error(),
		})
		return
	}

	db.Table("alunos").
		Joins("JOIN disciplina_matriculas dm ON dm.aluno_unique_id = alunos.unique_id").
		Joins("JOIN curso_disciplinas cd ON cd.unique_id = dm.curso_disciplina_unique_id").
		Joins("JOIN cursos c ON c.unique_id = cd.curso_unique_id").
		Joins("JOIN faculdades f ON f.unique_id = c.faculdade_unique_id").
		Where("f.unique_id = ?", newid).
		Distinct("alunos.unique_id").
		Count(&totalAlunos)

	c.JSON(200, totalAlunos)
}

// @Summary Busca aluno por id
// @Description Busca aluno por id
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do aluno"
// @Success 200 {object} models.Aluno
// @Router /aluno/{id} [GET]
func BuscarAluno(c *gin.Context) {
	var aluno models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&aluno, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "aluno não encontrado com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, aluno)
}

// @Summary Busca todos os alunos
// @Description Busca todos os alunos
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /aluno [GET]
func BuscarAlunos(c *gin.Context) {
	var alunos []models.Aluno
	db := database.GetDatabase()

	err := db.Find(&alunos).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar os alunos " + err.Error(),
		})
		return
	}

	c.JSON(200, alunos)
}

// @Summary Cria um aluno
// @Description Cria um aluno, os campos que devem ser enviados são: nome e cpf
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "aluno"
// @Success 201 {object} models.Aluno
// @Router /aluno [POST]
func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	aluno.Cpf = cpfcnpj.Clean(aluno.Cpf)
	resultadoValidador := cpfcnpj.ValidateCPF(aluno.Cpf)
    if !resultadoValidador {
		c.JSON(400, gin.H{
			"erro": "cpf inválido",
		})
		return
    }

	validate := validator.New()

    err = validate.Struct(aluno)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	err = db.Create(&aluno).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o aluno: " + err.Error(),
		})
		return
	}

	c.JSON(201, aluno)
}

// @Summary Deleta um aluno
// @Description Deleta um aluno
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do aluno"
// @Success 204 
// @Router /aluno/{id} [DELETE]
func DeletarAluno(c *gin.Context) {
	var aluno models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&aluno, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Aluno não encontrado: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.Aluno{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar o aluno: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita um aluno
// @Description Edita um aluno, os campos que devem ser enviados são: nome e cpf
// @Tags Controller Aluno
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do aluno"
// @Param aluno body models.Aluno true "aluno"
// @Success 200 {object} models.Aluno
// @Router /aluno/{id} [PUT]
func EditarAluno(c *gin.Context) {
	var aluno models.Aluno
	db := database.GetDatabase()
	id := c.Param("id")
	
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&aluno, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Aluno não encontrado: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&aluno)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	aluno.Cpf = cpfcnpj.Clean(aluno.Cpf)
	resultadoValidador := cpfcnpj.ValidateCPF(aluno.Cpf)
    if !resultadoValidador {
		c.JSON(400, gin.H{
			"erro": "cpf inválido",
		})
		return
    }

	validate := validator.New()

    err = validate.Struct(aluno)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	err = db.Save(&aluno).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o aluno: " + err.Error(),
		})
		return
	}

	c.JSON(200, aluno)
}

