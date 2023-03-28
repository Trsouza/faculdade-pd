package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
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
// @Param curso body models.Curso true "curso"
// @Success 201 {object} models.Curso
// @Router /curso [POST]
func CriarCurso(c *gin.Context) {
	var curso models.Curso
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&curso)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Create(&curso).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o curso: " + err.Error(),
		})
		return
	}

	c.JSON(201, curso)
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
// @Description Edita um curso, os campos que devem ser enviados são: nome e "faculdade_unique_id
// @Tags Controller Curso
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do curso"
// @Param curso body models.Curso true "curso"
// @Success 200 {object} models.Curso
// @Router /curso/{id} [PUT]
func EditarCurso(c *gin.Context) {
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

	err = c.ShouldBindJSON(&curso)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Save(&curso).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o curso: " + err.Error(),
		})
		return
	}

	c.JSON(200, curso)
}

