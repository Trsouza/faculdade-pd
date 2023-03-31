package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
	"github.com/trsouza/faculdade-pd/dtos"
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
// @Param cursoDiscipliona body dtos.CursoDisciplinaCreateDTO true "cursoDiscipliona"
// @Success 201 {object} models.CursoDisciplina
// @Router /curso-disciplina [POST]
func CriarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina dtos.CursoDisciplinaCreateDTO
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&cursoDisciplina)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	validate := validator.New()

    err = validate.Struct(cursoDisciplina)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	novoCursoDisciplina := models.CursoDisciplina{
		CursoUniqueID: cursoDisciplina.CursoUniqueID,
		DisciplinaUniqueID: cursoDisciplina.DisciplinaUniqueID,
	}

	err = db.Create(&novoCursoDisciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar o cursoDisciplina: " + err.Error(),
		})
		return
	}

	c.JSON(201, novoCursoDisciplina)
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
// @Description Edita um cursoDiscipliona, os campos que podem ser enviados são: curso_unique_id e "disciplina_unique_id
// @Tags Controller CursoDisciplina
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do cursoDiscipliona"
// @Param cursoDiscipliona body dtos.CursoDisciplinaUpdateDTO true "cursoDiscipliona"
// @Success 200 {object} models.CursoDisciplina
// @Router /curso-disciplina/{id} [PUT]
func EditarCursoDisciplina(c *gin.Context) {
	var cursoDisciplina dtos.CursoDisciplinaUpdateDTO
	var novoCursoDisciplina models.CursoDisciplina
	db := database.GetDatabase()
	id := c.Param("id")
	
	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&novoCursoDisciplina, newid).Error
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

	if cursoDisciplina.CursoUniqueID > 0 {  novoCursoDisciplina.CursoUniqueID = cursoDisciplina.CursoUniqueID  }
	if cursoDisciplina.DisciplinaUniqueID > 0 {  novoCursoDisciplina.DisciplinaUniqueID = cursoDisciplina.DisciplinaUniqueID  }

	err = db.Save(&novoCursoDisciplina).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar o cursoDisciplina: " + err.Error(),
		})
		return
	}

	c.JSON(200, novoCursoDisciplina)
}

