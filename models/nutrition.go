package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Nutrition struct {
		ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		CalorieMax int       `json:"calorie_max" gorm:"type:int;null" binding:"omitempty,min=1,max=9999"`
		FatMax     float64   `json:"fat_max" gorm:"type:numeric;null" binding:"omitempty,min=0,max=999"`
		ProteinMin float64   `json:"protein_min" gorm:"type:numeric;null" binding:"omitempty,min=0,max=999"`
		CarbMax    float64   `json:"carb_max" gorm:"type:numeric;null" binding:"omitempty,min=0,max=999"`
		SugarMax   float64   `json:"sugar_max" gorm:"type:numeric;null" binding:"omitempty,min=0,max=999"`
		SodiumMax  float64   `json:"sodium_max" gorm:"type:numeric;null" binding:"omitempty,min=0,max=9999"`
		FiberMin   float64   `json:"fiber_min" gorm:"type:numeric;null" binding:"omitempty,min=0,max=999"`
		CreatedAt  time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
