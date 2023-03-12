package book

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB used to init the remote DB (PostgreSQL as a Service)
func InitDB() *gorm.DB {
	// Databse info from ElephantDB (PostgreSQL as a Service Cloud Database)
	dsn := "host=mahmud.db.elephantsql.com user=xhnxbuqe password=kGHMnhnc-hbhVZY04JzZP1VUfv_7-OIh dbname=xhnxbuqe port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	return db
}
