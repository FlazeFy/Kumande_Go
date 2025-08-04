package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	CountCalorie struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		Weight    int       `json:"weight" gorm:"type:int;not null" binding:"required,min=1,max=999"`
		Height    int       `json:"height" gorm:"type:int;not null" binding:"required,min=1,max=999"`
		Result    int       `json:"result" gorm:"type:int;not null" binding:"required,min=1,max=99999"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
