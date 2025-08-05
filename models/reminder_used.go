package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	ReminderUsed struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		// FK - Reminder
		ReminderId uuid.UUID `json:"reminder_id" gorm:"not null" binding:"required,max=36,min=36"`
		Reminder   Reminder  `json:"-" gorm:"foreignKey:ReminderId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
