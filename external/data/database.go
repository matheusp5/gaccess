package data

import (
	"gaccess/internal/logger"
	"gaccess/logic/entities"
	"gorm.io/gorm"
)

var newLogger logger.Logger = logger.NewLogger()

func openDatabase() *gorm.DB {
	sqlite := OpenSqlite()
	db, err := gorm.Open(sqlite, &gorm.Config{})
	if err != nil {
		newLogger.Error("An error occurred while opening the dialect: " + err.Error())
		panic(nil)
	}
	newLogger.Info("Database opened with success.")
	return db
}

func migrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(&entities.SiteRedirect{}, &entities.AccessInfo{})
	if err != nil {
		newLogger.Error("An error occurred while migrating the database: " + err.Error())
		panic(nil)
	}
	newLogger.Info("Database migrated with success.")
}

func InitDatabase() *gorm.DB {
	db := openDatabase()
	migrateDatabase(db)
	return db
}
