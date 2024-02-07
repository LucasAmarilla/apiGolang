package controllers

import (
	"github.com/gin-gonic/gin"
	initializers "github.com/lucasamarilla/go-crud/initializer"
	"github.com/lucasamarilla/go-crud/models"
)

func MembroCreate(c *gin.Context) {

	var body struct {
		Name       string
		Nascimento string
		Foto       string
		Meme       string
		Desc       string
	}

	c.Bind(&body)

	membro := models.Membro{Name: body.Name, Nascimento: body.Nascimento, Foto: body.Foto, Meme: body.Meme, Desc: body.Desc}
	result := initializers.DB.Create(&membro)

	if result != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Membro": membro,
	})
}

func MembrosIndex(c *gin.Context) {
	var membros []models.Membro
	initializers.DB.Find(&membros)

	c.JSON(200, gin.H{
		"Membros": membros,
	})

}

func ShowMembro(c *gin.Context) {

	id := c.Param("id")

	var membro models.Membro
	initializers.DB.First(&membro, id)

	c.JSON(200, gin.H{
		"Membro": membro,
	})

}

func UpdateMembro(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name       string
		Nascimento string
		Foto       string
		Meme       string
		Desc       string
	}

	c.Bind(&body)

	var membro models.Membro
	initializers.DB.First(&membro, id)

	initializers.DB.Model(&membro).Updates(models.Membro{
		Name:       body.Name,
		Nascimento: body.Nascimento,
		Foto:       body.Foto,
		Meme:       body.Meme,
		Desc:       body.Desc,
	})

	c.JSON(200, gin.H{
		"Membro": membro,
	})
}

func DeleteMembro(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Membro{}, id)
	c.Status(200)
}
