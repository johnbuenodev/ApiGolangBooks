package controllers

import (
	"apiGolang/database"
	"apiGolang/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
func TestEndPointBook(c *gin.Context) {

	c.JSON(200, gin.H{
		"valeu": "ok",
	})
} */

func ShowBook(c *gin.Context) {

	//c.JSON(200, gin.H{
	//	"valeu": "ok",
	//})

	//Pega o Parametro Id pelo C de Contexto

	id := c.Param("id")

	//Pega o id que vem como string e convert para inteiro int
	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()

	var book models.Book
	err = db.First(&book, newid).Error

	//ifer autocomplete
	//Se não for encontrado retorna Erro
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi encontrado Book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {

	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel realizar o bind do Book: " + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel criar o Book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDatabase()

	//Array de books para retornar todos os livros
	var books []models.Book
	err := db.Find(&books).Error

	if err != nil {

		c.JSON(400, gin.H{
			"Error": "Não foi possivel localizar os Books: " + err.Error(),
		})
		return
	}

	c.JSON(200, books)

}

//Passa todo o objeto o Save do Database fica responsavel por atualizar o registro caso ele exista
func UpdateBook(c *gin.Context) {

	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel realizar o bind do Book: " + err.Error(),
		})
		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel Atualizar o Book: " + err.Error(),
		})
		return
	}

	c.JSON(200, book)

}

func DeleteBook(c *gin.Context) {

	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID precisa ser um inteiro",
		})
		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, newid).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Não foi possivel Deletar o Book: " + err.Error(),
		})
		return
	}

	//Retorna somente status 204
	//No Content indica que a solicitação foi bem sucedida
	c.Status(204)
}
