package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func GetDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect Database")
	}
	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
	seeder(db)
	return db
}
