package config

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	"github.com/heroku/go-getting-started/controllers"
)

// Connecting to db
func Connect() *pg.DB {

	// Connecting to db
	opts := &pg.Options{
		User: "wdeoibbwbhscjy",
		Password: "6931402f0ef2741c19a6ed4d196a7ec27662c2233cce0d3f6d072c154609629d",
		Addr: "postgres://wdeoibbwbhscjy:6931402f0ef2741c19a6ed4d196a7ec27662c2233cce0d3f6d072c154609629d@ec2-107-23-41-227.compute-1.amazonaws.com:5432/dg9keg1a0efbf",
		Database: "dg9keg1a0efbf",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	}
	log.Printf("Connected to db")
	
	



	controllers.CreateExerciseTable(db)
	controllers.InitiateDB(db)


	return db
}