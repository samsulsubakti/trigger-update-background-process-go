// db.go
package db

import (
	"log"
	"os"
	"sourceanddestination/controller"
	"sourceanddestination/migrations"
	"sourceanddestination/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// DB SOURCE
	dbnameSource := os.Getenv("DB_SOURCE")
	dbSource, errSource := gorm.Open(mysql.Open(dbnameSource), &gorm.Config{})

	if errSource != nil {
		panic("failed to connect database db_source")
	}

	log.Println("Database db_source connected...")
	migrations.RunMigration(dbSource, models.SourceProduct{})

	// DB DESTINATION
	dbnameDestination := os.Getenv("DB_DESTINATION")
	dbDestination, errDestination := gorm.Open(mysql.Open(dbnameDestination), &gorm.Config{})

	if errDestination != nil {
		panic("failed to connect database db_destination")
	}

	log.Println("Database db_destination connected...")
	migrations.RunMigration(dbDestination, models.DestinationProduct{})
	migrations.GenerateSeedData(dbSource, dbDestination)
	controller.Init(dbSource, dbDestination)
}
