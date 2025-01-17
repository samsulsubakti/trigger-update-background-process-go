package migrations

import (
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB, model interface{}) {
	tx := db.Begin()

	db.AutoMigrate(&model)

	tx.Commit()
}
