package data

import (
	"gaccess/internal/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const SQLITE_FILE string = "gaccess.db"

func OpenSqlite() gorm.Dialector {
	db := sqlite.Open(SQLITE_FILE)
	newLogger := logger.NewLogger()
	newLogger.Info("SQLite database opened with success.")
	return db
}
