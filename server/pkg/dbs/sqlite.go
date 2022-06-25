package dbs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSqliteDB(entities []interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	for _, entity := range entities {
		err = db.AutoMigrate(entity)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
