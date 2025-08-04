package modules

import (
	"kumande/models"
	"log"

	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Admin{},
		&models.History{},
		&models.Error{},
		&models.Feedback{},
		&models.AuditTrail{},
		&models.Dictionary{},
		&models.Allergic{},
		&models.BodyInfo{},
		&models.Budget{},
		&models.Tag{},
		&models.Reminder{},
		&models.UserTrack{},
		&models.UserWeather{},
		&models.Consume{},
	)

	if err != nil {
		panic(err.Error())
	}

	log.Println("Migrate Success!")
}
