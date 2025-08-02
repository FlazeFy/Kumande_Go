package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	BodyInfo struct {
		ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		BloodPressure string    `json:"blood_pressure" gorm:"type:varchar(9);not null" binding:"required,min=6,max=9"`
		BloodGlucose  int       `json:"blood_glucose" gorm:"type:int;not null" binding:"required,max=3"`
		Gout          float32   `json:"gout" gorm:"type:double precision;not null" binding:"required"`
		Cholesterol   int       `json:"cholesterol" gorm:"type:int;not null" binding:"required,max=3"`
		CreatedAt     time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
