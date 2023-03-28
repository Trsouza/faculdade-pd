package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
)

// @Summary Busca professor por id
// @Description Busca professor por id
// @Tags Controller Professor
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do professor"
// @Success 200 {object} models.Professor
// @Router /professor/{id} [GET]
func BuscarProfessor(c *gin.Context) {
	var professor models.Professor
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&professor, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "professor não encontrado com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, professor)
}

// @Summary Buscar todos os professors
// @Description Busca todos os professors
// @Tags Controller Professor
// @Accept json
// @Produce json
// @Success 200 {object} models.Professor
// @Router /professor [GET]
func BuscarProfessores(c *gin.Context) {
	var professor []models.Professor
	db := database.GetDatabase()

	err := db.Find(&professor).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar os professores " + err.Error(),
		})
		return
	}

	c.JSON(200, professor)
}

// @Summary Cria um professor
// @Description Cria um professor, os campos que devem ser enviados são: nome e formacao
// @Tags Controller Professor
// @Accept json
// @Produce json
// @Param professor body models.Professor true "professor"
// @Success 201 {object} models.Professor
// @Router /professor [POST]
func CriarProfessor(c *gin.Context) {
	var professor models.Professor
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&professor)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	validate := validator.New()

    err = validate.Struct(professor)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	err = db.Create(&professor).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o professor: " + err.Error(),
		})
		return
	}

	c.JSON(201, professor)
}

// @Summary Deleta um professor
// @Description Deleta um professor
// @Tags Controller Professor
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do professor"
// @Success 204 
// @Router /professor/{id} [DELETE]
func DeletarProfessor(c *gin.Context) {
	var professor models.Professor
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&professor, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Professor não encontrado: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.Professor{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar o professor: " + err.Error(),
		})
		return
	}

	c.Status(204)
}

// @Summary Edita um professor
// @Description Edita um professor, os campos que devem ser enviados são: nome e formacao
// @Tags Controller Professor
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do professor"
// @Param professor body models.Professor true "professor"
// @Success 200 {object} models.Professor
// @Router /professor/{id} [PUT]
func EditarProfessor(c *gin.Context) {
	var professor models.Professor
	db := database.GetDatabase()
	id := c.Param("id")
	
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&professor, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Professor não encontrado: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&professor)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	validate := validator.New()

    err = validate.Struct(professor)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	err = db.Save(&professor).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o professor: " + err.Error(),
		})
		return
	}

	c.JSON(200, professor)
}

