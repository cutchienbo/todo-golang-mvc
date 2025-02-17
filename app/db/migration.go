package db

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"golang-mvc/app/models/db"
)

func Migrate(gormDB *gorm.DB) {
	// Define migrations
	m := gormigrate.New(gormDB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "2025021701",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&db.User{}, &db.Todo{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&db.Todo{}, &db.User{})
			},
		},
	})

	if err := m.Migrate(); err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	println("Migrations applied successfully")
}