package book

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// databse info from ElephantDB
	// postgres://xhnxbuqe:kGHMnhnc-hbhVZY04JzZP1VUfv_7-OIh@mahmud.db.elephantsql.com/xhnxbuqe

	dsn := "host=mahmud.db.elephantsql.com user=xhnxbuqe password=kGHMnhnc-hbhVZY04JzZP1VUfv_7-OIh dbname=xhnxbuqe port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to Database")
	}

	return db
}
