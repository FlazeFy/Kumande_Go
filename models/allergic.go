package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Allergic struct {
		ID              uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
		AllergicContext string     `json:"allergic_context" gorm:"type:varchar(75);not null" binding:"required,min=1,max=75"`
		AllergicDesc    *string    `json:"allergic_desc" gorm:"type:varchar(255);null" binding:"omitempty,min=1,max=255"`
		CreatedAt       time.Time  `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt       *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
