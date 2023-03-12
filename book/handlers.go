package book

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/selvan12/go-gin-gorm-postgres/model"
)

var bookEntry struct {
	Title string
	Price string
}

func (a *App) createBook(c *gin.Context) {
	// get the request body
	c.BindJSON(&bookEntry)

	// insert the data
	entry := model.Book{Title: bookEntry.Title, Price: bookEntry.Price}
	result := a.DB.Create(&entry)
	if result.Error != nil {
		log.Fatal("Failed to create an entry in book db")
		c.Status(400)
		return
	}

	// respond with success data
	c.JSON(201, entry)
}

func (a *App) listBook(c *gin.Context) {
	// get all data from database
	books := []model.Book{}
	result := a.DB.Find(&books)
	if result.Error != nil {
		log.Fatal("Failed to retrive the book entries from db")
		c.Status(400)
		return
	}

	// respond with results
	c.JSON(200, gin.H{
		"books": books,
	})
}

func (a *App) getBook(c *gin.Context) {
	// get a id of the specific book to retreive
	id := c.Param("id")

	// get specific data from database
	book := model.Book{}
	result := a.DB.First(&book, id)
	if result.Error != nil {
		log.Fatal("Failed to retrive the book entry from db")
		c.Status(400)
		return
	}

	// respond with results
	c.JSON(200, gin.H{
		"book": book,
	})
}

func (a *App) updateBook(c *gin.Context) {
	// get a id of the specific book to retreive
	id := c.Param("id")

	// get the request body
	c.BindJSON(&bookEntry)

	// get specific data from database
	book := model.Book{}
	result := a.DB.First(&book, id)
	if result.Error != nil {
		log.Fatal("Failed to retrive the book entry from db")
		c.Status(400)
		return
	}

	// update the book values
	result = a.DB.Model(&book).Updates(&model.Book{Title: bookEntry.Title, Price: bookEntry.Price})
	if result.Error != nil {
		log.Fatal("Failed to update the book entry in db")
		c.Status(400)
		return
	}

	// respond with results
	c.JSON(200, gin.H{
		"book": book,
	})
}

func (a *App) deleteBook(c *gin.Context) {
	// get a id of the specific book to retreive
	id := c.Param("id")

	// delete the specific book entry from database
	result := a.DB.Delete(&model.Book{}, id)
	if result.Error != nil {
		log.Fatal("Failed to delete the book entry in db")
		c.Status(400)
		return
	}

	// respond with success
	c.Status(204)
}
