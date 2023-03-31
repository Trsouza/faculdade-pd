package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
	"github.com/trsouza/faculdade-pd/dtos"
)

// @Summary Busca curso por id
// @Description Busca curso por id
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do curso"
// @Success 200 {object} models.Curso
// @Router /curso/{id} [GET]
func BuscarCurso(c *gin.Context) {
	var curso models.Curso
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

	c.JSON(200, curso)
}

// @Summary Buscar todos os cursos
// @Description Busca todos os cursos
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Success 200 {object} models.Curso
// @Router /curso [GET]
func BuscarCursos(c *gin.Context) {
	var curso []models.Curso
	db := database.GetDatabase()

	err := db.Find(&curso).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar os cursos " + err.Error(),
		})
		return
	}

	c.JSON(200, curso)
}

// @Summary Cria um curso
// @Description Cria um curso, os campos que devem ser enviados são: nome e "faculdade_unique_id
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Param curso body dtos.CursoCreateDTO true "curso"
// @Success 201 {object} models.Curso
// @Router /curso [POST]
func CriarCurso(c *gin.Context) {
	var curso dtos.CursoCreateDTO
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&curso)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	validate := validator.New()

    err = validate.Struct(curso)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	novoCurso := models.Curso{
		Nome: curso.Nome,
		FaculdadeUniqueID: curso.FaculdadeUniqueID,
	}

	err = db.Create(&novoCurso).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o curso: " + err.Error(),
		})
		return
	}

	c.JSON(201, novoCurso)
}

// @Summary Deleta um curso
// @Description Deleta um curso
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do curso"
// @Success 204 
// @Router /curso/{id} [DELETE]
func DeletarCurso(c *gin.Context) {
	var curso models.Curso
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
		c.JSON(404, gin.H{
			"erro": "Curso não encontrado: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.Curso{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar o curso: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita um curso
// @Description Edita um curso, os campos que podem ser enviados são: nome e "faculdade_unique_id
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do curso"
// @Param curso body dtos.CursoUpdateDTO true "curso"
// @Success 200 {object} models.Curso
// @Router /curso/{id} [PUT]
func EditarCurso(c *gin.Context) {
	var curso dtos.CursoUpdateDTO
	var novoCurso models.Curso
	db := database.GetDatabase()
	id := c.Param("id")
	
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&novoCurso, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Curso não encontrado: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&curso)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	if curso.Nome != "" {  novoCurso.Nome = curso.Nome  }
	if curso.FaculdadeUniqueID > 0 {  novoCurso.FaculdadeUniqueID = curso.FaculdadeUniqueID  }

	err = db.Save(&novoCurso).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o curso: " + err.Error(),
		})
		return
	}

	c.JSON(200, novoCurso)
}

