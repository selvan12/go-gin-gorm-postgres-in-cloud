package main

import (
	"log"

	"github.com/selvan12/go-gin-gorm-postgres/book"
)

func main() {
	log.Println("Hello welcome to go-gin-gorm-postgres project")

	app := book.NewApp()
	app.AddRoutes(app.Engine)

	//r := gin.Default()
	// app.Engine.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	app.Engine.Run() // listen and serve on 0.0.0.0:8080
}
