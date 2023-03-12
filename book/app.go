package book

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/selvan12/go-gin-gorm-postgres/model"
	"gorm.io/gorm"
)

type App struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func NewApp() *App {
	log.Println("App In")

	db := InitDB()
	db.AutoMigrate(&model.Book{})
	app := App{
		Engine: gin.Default(),
		DB:     db,
	}
	return &app
}
