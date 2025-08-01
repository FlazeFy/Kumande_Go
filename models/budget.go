package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Budget struct {
		ID          uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
		BudgetTotal string     `json:"budget_total" gorm:"type:int;not null" binding:"required,min=1,max=99999999"`
		BudgetMonth string     `json:"budget_month" gorm:"type:varchar(3);not null" binding:"required,max=3"`
		BudgetYear  int        `json:"budget_year" gorm:"type:int;not null" binding:"required,max=4"`
		CreatedAt   time.Time  `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt   *time.Time `json:"updated_at" gorm:"type:timestamp;null"`
		OverAt      *time.Time `json:"over_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
