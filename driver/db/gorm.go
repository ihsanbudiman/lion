package db

import (
	"lion/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(conf config.DatabaseConfig) *gorm.DB {

	db, err := gorm.Open(postgres.Open(conf.DSN()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
