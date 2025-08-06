package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	ConsumeRate struct {
		ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		ConsumeComment *string   `json:"consume_comment" gorm:"type:varchar(255);null" binding:"omitempty,max=255"`
		ConsumeRate    int       `json:"consume_rate" gorm:"type:int;not null" binding:"required,min=1,max=10"`
		CreatedAt      time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		// FK - Consume
		ConsumeId uuid.UUID `json:"consume_id" gorm:"not null" binding:"required,max=36,min=36"`
		Consume   Consume   `json:"-" gorm:"foreignKey:ConsumeId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
