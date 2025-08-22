package config

import (
	"os"

	"github.com/ImHypna/gopportunities.git/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//check if database file exists

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating the database...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.ErrorF("Sqlite opening error: %v", err)

		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrorF("sqlite automigration error: %v", err)
		return nil, err

	}

	return db, nil

}
