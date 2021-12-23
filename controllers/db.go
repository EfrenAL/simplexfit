package controllers

import (
	"database/sql"

	"gorm.io/gorm"
)


var dbConnect *sql.DB
var gormDBConnect *gorm.DB

func InitiateDB(db *sql.DB, gormDB *gorm.DB ) {
	dbConnect = db
	gormDBConnect = gormDB
}