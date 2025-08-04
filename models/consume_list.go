package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	ConsumeList struct {
		ID                uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		ConsumeListName   string         `json:"consume_list_name" gorm:"type:varchar(75);not null" binding:"required,max=75"`
		ConsumeListDetail *string        `json:"consume_list_detail" gorm:"type:varchar(255);null" binding:"omitempty,max=255"`
		ConsumeListTag    datatypes.JSON `json:"consume_list_tag" gorm:"type:json;null" binding:"omitempty"`
		CreatedAt         time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt         *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
