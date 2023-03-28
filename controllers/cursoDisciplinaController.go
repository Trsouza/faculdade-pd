package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
)

// @Summary Busca cursoDisciplina por id
// @Description Busca cursoDisciplina por id
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do cursoDisciplina"
// @Success 200 {object} models.CursoDisciplina
// @Router /curso-disciplina/{id} [GET]
func BuscarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina models.CursoDisciplina
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&cursoDisciplina, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "cursoDisciplina não encontrado com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, cursoDisciplina)
}

// @Summary Buscar todos os cursoDisciplionas
// @Description Busca todos os cursoDisciplinas
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Success 200 {object} models.CursoDisciplina
// @Router /curso-disciplina [GET]
func BuscarCursoDisciplinas(c *gin.Context) {
	var cursoDisciplina []models.CursoDisciplina
	db := database.GetDatabase()

	err := db.Find(&cursoDisciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar os cursoDisciplinas " + err.Error(),
		})
		return
	}

	c.JSON(200, cursoDisciplina)
}

// @Summary Cria um cursoDisciplionas
// @Description Cria um cursoDisciplionas
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Param cursoDiscipliona body models.CursoDisciplina true "cursoDiscipliona"
// @Success 201 {object} models.CursoDisciplina
// @Router /curso-disciplina [POST]
func CriarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina models.CursoDisciplina
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&cursoDisciplina)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Create(&cursoDisciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o cursoDisciplina: " + err.Error(),
		})
		return
	}

	c.JSON(201, cursoDisciplina)
}

// @Summary Deleta um cursoDiscipliona
// @Description Deleta um cursoDiscipliona, os campos que devem ser enviados são: curso_unique_id e "disciplina_unique_id
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do cursoDiscipliona"
// @Success 204 
// @Router /curso-disciplina/{id} [DELETE]
func DeletarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina models.CursoDisciplina
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&cursoDisciplina, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "CursoDisciplina não encontrado: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.CursoDisciplina{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar o cursoDisciplina: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita um cursoDiscipliona
// @Description Edita um cursoDiscipliona, os campos que devem ser enviados são: curso_unique_id e "disciplina_unique_id
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do cursoDiscipliona"
// @Param cursoDiscipliona body models.CursoDisciplina true "cursoDiscipliona"
// @Success 200 {object} models.CursoDisciplina
// @Router /curso-disciplina/{id} [PUT]
func EditarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina models.CursoDisciplina
	db := database.GetDatabase()
	id := c.Param("id")
	
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&cursoDisciplina, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "CursoDisciplina não encontrado: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&cursoDisciplina)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	err = db.Save(&cursoDisciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o cursoDisciplina: " + err.Error(),
		})
		return
	}

	c.JSON(200, cursoDisciplina)
}

