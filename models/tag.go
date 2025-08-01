package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Tag struct {
		ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
		TagSlug   string    `json:"tag_slug" gorm:"type:varchar(46);not null" binding:"required,min=1,max=46"`
		TagName   string    `json:"tag_name" gorm:"type:varchar(36);not null" binding:"required,min=1,max=36"`
		CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;not null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)
