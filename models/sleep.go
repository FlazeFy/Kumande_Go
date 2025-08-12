package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Sleep struct {
		ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
		SleepNote    *string    `json:"sleep_note" gorm:"type:varchar(255);null" binding:"omitempty,max=255"`
		SleepQuality int        `json:"sleep_quality" gorm:"type:int;not null" binding:"omitempty,min=1,max=10"`
		CreatedAt    time.Time  `json:"created_at" gorm:"type:timestamp;not null"`
		SleepAt      time.Time  `json:"sleep_at" gorm:"type:timestamp;not null"`
		WokeAt       *time.Time `json:"woke_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
