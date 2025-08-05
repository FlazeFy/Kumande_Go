package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	ConsumeListRelation struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		// FK - Consume
		ConsumeId uuid.UUID `json:"consume_id" gorm:"not null" binding:"required,max=36,min=36"`
		Consume   Consume   `json:"-" gorm:"foreignKey:ConsumeId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		// FK - Consume List
		ConsumeListId uuid.UUID   `json:"consume_list_id" gorm:"not null" binding:"required,max=36,min=36"`
		ConsumeList   ConsumeList `json:"-" gorm:"foreignKey:ConsumeListId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
