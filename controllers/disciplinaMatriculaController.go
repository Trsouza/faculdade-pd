package controllers

import (
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
	"github.com/trsouza/faculdade-pd/dtos"
)

// @Summary Busca disciplinaMatricula por id
// @Description Busca disciplinaMatricula por id
// @Tags Controller DisciplinaMatricula
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do disciplinaMatricula"
// @Success 200 {object} models.DisciplinaMatricula
// @Router /disciplina-matricula/{id} [GET]
func BuscarDisciplinaMatricula(c *gin.Context) {
	var disciplinaMatricula models.DisciplinaMatricula
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&disciplinaMatricula, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "disciplinaMatricula não encontrada com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, disciplinaMatricula)
}

// @Summary Buscar todas as disciplinas
// @Description Busca todas as disciplinas
// @Tags Controller DisciplinaMatricula
// @Accept json
// @Produce json
// @Success 200 {object} models.DisciplinaMatricula
// @Router /disciplina-matricula [GET]
func BuscarDisciplinaMatriculas(c *gin.Context) {
	var disciplinaMatricula []models.DisciplinaMatricula
	db := database.GetDatabase()

	err := db.Find(&disciplinaMatricula).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar as disciplinaMatriculas " + err.Error(),
		})
		return
	}

	c.JSON(200, disciplinaMatricula)
}

// @Summary Cria uma disciplinaMatricula
// @Description Cria uma disciplinaMatricula, os campos que devem ser enviados são: aluno_unique_id e curso_disciplina_unique_id
// @Tags Controller DisciplinaMatricula
// @Accept json
// @Produce json
// @Param disciplinaMatricula body dtos.DisciplinaMatriculaCreateDTO true "disciplinaMatricula"
// @Success 201 {object} models.DisciplinaMatricula
// @Router /disciplina-matricula [POST]
func CriarDisciplinaMatricula(c *gin.Context) {
	var disciplinaMatricula dtos.DisciplinaMatriculaCreateDTO
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&disciplinaMatricula)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	
	validate := validator.New()

    err = validate.Struct(disciplinaMatricula)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	novaDisciplinaMatricula := models.DisciplinaMatricula{
		CursoDisciplinaUniqueID: disciplinaMatricula.CursoDisciplinaUniqueID,
		AlunoUniqueID: disciplinaMatricula.AlunoUniqueID,
		DataMatricula: time.Now(),
	}

	err = db.Create(&novaDisciplinaMatricula).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar a disciplinaMatricula: " + err.Error(),
		})
		return
	}

	c.JSON(201, novaDisciplinaMatricula)
}

// @Summary Deleta uma disciplinaMatricula
// @Description Deleta uma disciplinaMatricula
// @Tags Controller DisciplinaMatricula
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da disciplinaMatricula"
// @Success 204 
// @Router /disciplina-matricula/{id} [DELETE]
func DeletarDisciplinaMatricula(c *gin.Context) {
	var disciplinaMatricula models.DisciplinaMatricula
	id := c.Param("id")
	db := database.GetDatabase()

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&disciplinaMatricula, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "DisciplinaMatricula não encontrada: " + err.Error(),
		})
		return
	}
	
	err = db.Delete(&models.DisciplinaMatricula{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar a disciplinaMatricula: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita uma disciplinaMatricula
// @Description Edita uma disciplinaMatricula, os campos que podem ser enviados são: aluno_unique_id e curso_disciplina_unique_id
// @Tags Controller DisciplinaMatricula
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da disciplinaMatricula"
// @Param disciplinaMatricula body dtos.DisciplinaMatriculaUpdateDTO true "disciplinaMatricula"
// @Success 200 {object} models.DisciplinaMatricula
// @Router /disciplina-matricula/{id} [PUT]
func EditarDisciplinaMatricula(c *gin.Context) {
	var disciplinaMatricula dtos.DisciplinaMatriculaUpdateDTO
	var novaDisciplinaMatricula models.DisciplinaMatricula
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&novaDisciplinaMatricula, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "DisciplinaMatricula não encontrada: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&disciplinaMatricula)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	if disciplinaMatricula.CursoDisciplinaUniqueID > 0 { novaDisciplinaMatricula.CursoDisciplinaUniqueID = disciplinaMatricula.CursoDisciplinaUniqueID}
	if disciplinaMatricula.AlunoUniqueID > 0 { novaDisciplinaMatricula.AlunoUniqueID = disciplinaMatricula.AlunoUniqueID}

	err = db.Save(&novaDisciplinaMatricula).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar a disciplinaMatricula: " + err.Error(),
		})
		return
	}

	c.JSON(200, novaDisciplinaMatricula)
}

