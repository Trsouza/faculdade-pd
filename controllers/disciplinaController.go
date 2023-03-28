package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
)

// @Summary Filtra disciplinas que um determinado aluno cursa
// @Description Realiza um filtro para obter quais disciplinas um determinado aluno cursa
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do aluno"
// @Success 200 {object} models.Disciplina
// @Router /disciplina/aluno/{id} [GET]
func FiltrarDisciplinasPorAluno(c *gin.Context) {
	var aluno models.Aluno
	var disciplinas []models.Disciplina
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

	db.Table("disciplinas").
		Joins("JOIN curso_disciplinas cd ON cd.disciplina_unique_id = disciplinas.unique_id").
		Joins("JOIN disciplina_matriculas dm ON dm.curso_disciplina_unique_id = cd.unique_id").
		Joins("JOIN alunos a ON dm.aluno_unique_id = a.unique_id").
		Where("a.unique_id = ?", newid).
		Select("disciplinas.unique_id, disciplinas.nome").
		Find(&disciplinas)

	c.JSON(200, disciplinas)
}

// @Summary Busca disciplina por id
// @Description Busca disciplina por id
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do disciplina"
// @Success 200 {object} models.Disciplina
// @Router /disciplina/{id} [GET]
func BuscarDisciplina(c *gin.Context) {
	var disciplina models.Disciplina
	id := c.Param("id")
	db := database.GetDatabase()

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
			"erro": "disciplina não encontrada com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, disciplina)
}

// @Summary Buscar todas as disciplinas
// @Description Busca todas as disciplinas
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Success 200 {object} models.Disciplina
// @Router /disciplina [GET]
func BuscarDisciplinas(c *gin.Context) {
	var disciplina []models.Disciplina
	db := database.GetDatabase()

	err := db.Find(&disciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar as disciplinas " + err.Error(),
		})
		return
	}

	c.JSON(200, disciplina)
}

// @Summary Cria uma disciplina
// @Description Cria uma disciplina, os campos que devem ser enviados são: nome e professor_unique_id
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Param disciplina body models.Disciplina true "disciplina"
// @Success 201 {object} models.Disciplina
// @Router /disciplina [POST]
func CriarDisciplina(c *gin.Context) {
	var disciplina models.Disciplina
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&disciplina)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Create(&disciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar a disciplina: " + err.Error(),
		})
		return
	}

	c.JSON(201, disciplina)
}

// @Summary Deleta uma disciplina
// @Description Deleta uma disciplina
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da disciplina"
// @Success 204 
// @Router /disciplina/{id} [DELETE]
func DeletarDisciplina(c *gin.Context) {
	var disciplina models.Disciplina
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
		c.JSON(404, gin.H{
			"erro": "Disciplina não encontrada: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.Disciplina{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar a disciplina: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita uma disciplina
// @Description Edita uma disciplina, os campos que devem ser enviados são: nome e professor_unique_id
// @Tags Controller Disciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da disciplina"
// @Param disciplina body models.Disciplina true "disciplina"
// @Success 200 {object} models.Disciplina
// @Router /disciplina/{id} [PUT]
func EditarDisciplina(c *gin.Context) {
	var disciplina models.Disciplina
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
		c.JSON(404, gin.H{
			"erro": "Disciplina não encontrada: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&disciplina)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Save(&disciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar a disciplina: " + err.Error(),
		})
		return
	}

	c.JSON(200, disciplina)
}

