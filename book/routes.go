package book

import "github.com/gin-gonic/gin"

// AddRoutes helps to add all CRUD handler routes
func (a *App) AddRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// CRUD handlers for book
	r.POST("/book", a.createBook)
	r.GET("/book", a.listBook)
	r.GET("/book/:id", a.getBook)
	r.PATCH("/book/:id", a.updateBook)
	r.DELETE("/book/:id", a.deleteBook)
}
