package controllers

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
	"github.com/trsouza/faculdade-pd/database"
	"github.com/trsouza/faculdade-pd/models"
	"github.com/trsouza/faculdade-pd/dtos"
)

// @Summary Busca faculdade por id
// @Description Busca faculdade por id
// @Tags Controller Faculdade
// @Accept json
// @Produce json
// @Param id path int true "UniqueID do faculdade"
// @Success 200 {object} models.Faculdade
// @Router /faculdade/{id} [GET]
func BuscarFaculdade(c *gin.Context) {
	var faculdade models.Faculdade
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.Preload("Curso").First(&faculdade, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "faculdade não encontrada com o id: " + err.Error(),
		})
		return
	}

	c.JSON(200, faculdade)
}

// @Summary Buscar todas as faculdades
// @Description Busca todas as faculdades
// @Tags Controller Faculdade
// @Accept json
// @Produce json
// @Success 200 {object} models.Faculdade
// @Router /faculdade [GET]
func BuscarFaculdades(c *gin.Context) {
	db := database.GetDatabase()
	var faculdade []models.Faculdade
	err := db.Find(&faculdade).Error

	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível carregar as faculdades " + err.Error(),
		})
		return
	}

	c.JSON(200, faculdade)
}

// @Summary Cria uma faculdade
// @Description Cria uma faculdade, os campos que devem ser enviados são: nome e cnpj
// @Tags Controller Faculdade
// @Accept json
// @Produce json
// @Param faculdade body dtos.FaculdadeCreateDTO true "faculdade"
// @Success 201 {object} models.Faculdade
// @Router /faculdade [POST]
func CriarFaculdade(c *gin.Context) {
	var faculdade dtos.FaculdadeCreateDTO
	db := database.GetDatabase()

	err := c.ShouldBindJSON(&faculdade)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	validate := validator.New()

    err = validate.Struct(faculdade)
    if err != nil {
        c.JSON(400, gin.H{
			"erro": "objeto inválido: " + err.Error(),
		})
		return
    }

	
	faculdade.Cnpj = cpfcnpj.Clean(faculdade.Cnpj)
	resultadoValidador := cpfcnpj.ValidateCNPJ(faculdade.Cnpj)
    if !resultadoValidador {
		c.JSON(400, gin.H{
			"erro": "cnpj inválido",
		})
		return
    }

	novaFaculdade := models.Faculdade{
		Nome: faculdade.Nome,
		Cnpj:  faculdade.Cnpj,
	}

	err = db.Create(&novaFaculdade).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível salvar a faculdade: " + err.Error(),
		})
		return
	}

	c.JSON(201, novaFaculdade)
}

// @Summary Deleta uma faculdade
// @Description Deleta uma faculdade
// @Tags Controller Faculdade
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da faculdade"
// @Success 204 
// @Router /faculdade/{id} [DELETE]
func DeletarFaculdade(c *gin.Context) {
	var faculdade models.Faculdade
	id := c.Param("id")
	db := database.GetDatabase()

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&faculdade, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Faculdade não encontrada: " + err.Error(),
		})
		return
	}

	err = db.Delete(&models.Faculdade{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível deletar a faculdade: " + err.Error(),
		})
		return
	}

	c.Status(204)
}


// @Summary Edita uma faculdade
// @Description Edita uma faculdade, os campos que devem ser enviados são: nome e cnpj
// @Tags Controller Faculdade
// @Accept json
// @Produce json
// @Param id path int true "UniqueID da faculdade"
// @Param faculdade body dtos.FaculdadeUpdateDTO true "faculdade"
// @Success 200 {object} models.Faculdade
// @Router /faculdade/{id} [PUT]
func EditarFaculdade(c *gin.Context) {
	var faculdade dtos.FaculdadeUpdateDTO
	var novaFaculdade models.Faculdade
	db := database.GetDatabase()
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "ID deve ser um número inteiro",
		})
		return
	}

	err = db.First(&novaFaculdade, newid).Error
	if err != nil {
		c.JSON(404, gin.H{
			"erro": "Faculdade não encontrada: " + err.Error(),
		})
		return
	}

	err = c.ShouldBindJSON(&faculdade)
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "JSON inválido: " + err.Error(),
		})
		return
	}

	if faculdade.Cnpj != "" {
		faculdade.Cnpj = cpfcnpj.Clean(faculdade.Cnpj)
		resultadoValidador := cpfcnpj.ValidateCNPJ(faculdade.Cnpj)
		if !resultadoValidador {
			c.JSON(400, gin.H{
				"erro": "cnpj inválido",
			})
			return
		}
	}

	if faculdade.Nome != "" {  novaFaculdade.Nome = faculdade.Nome  }
	if faculdade.Cnpj != "" {  novaFaculdade.Cnpj = faculdade.Cnpj  }

	err = db.Save(&novaFaculdade).Error
	if err != nil {
		c.JSON(400, gin.H{
			"erro": "não foi possível editar a faculdade: " + err.Error(),
		})
		return
	}

	c.JSON(200, novaFaculdade)
}

