package controllers

import (
	"database/sql"

	"gorm.io/gorm"
)

//var dbConnect *sql.DB
var gormDBConnect *gorm.DB

func InitiateDB(db *sql.DB, gormDB *gorm.DB) {
	//dbConnect = db
	gormDBConnect = gormDB
}

func CreateTables() {

	gormDBConnect.Migrator().DropTable(&Exercise{}, &Workout{}, &User{})
	gormDBConnect.AutoMigrate(&Exercise{}, &Workout{}, &User{})

	gormDBConnect.Migrator().CreateTable(&Exercise{})
	gormDBConnect.Migrator().CreateTable(&Workout{})
	gormDBConnect.Migrator().CreateTable(&User{})

}
