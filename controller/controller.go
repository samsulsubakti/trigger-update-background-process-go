package controller

import "gorm.io/gorm"

var dbSource *gorm.DB
var dbDestination *gorm.DB

func Init(dbConnSorce *gorm.DB, dbConnDestination *gorm.DB) {
	dbSource = dbConnSorce
	dbDestination = dbConnDestination
}
