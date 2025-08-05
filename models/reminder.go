package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Reminder struct {
		ID                 uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		ReminderName       string         `json:"reminder_name" gorm:"type:varchar(75);not null" binding:"required,min=1,max=75"`
		ReminderType       string         `json:"reminder_type" gorm:"type:varchar(36);not null" binding:"required,min=1,max=36"`
		ReminderContext    datatypes.JSON `json:"reminder_context" gorm:"type:json;not null" binding:"required"`
		ReminderBody       string         `json:"reminder_body" gorm:"type:varchar(255);not null" binding:"required"`
		ReminderAttachment datatypes.JSON `json:"reminder_attachment" gorm:"type:json;null" binding:"omitempty"`
		CreatedAt          time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt          *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
	ReminderContext struct {
		Time string `json:"time"`
	}
	ReminderAttachment struct {
		AttachmentType    string `json:"attachment_type"`
		AttachmentContext string `json:"attachment_context"`
		AttachmentName    string `json:"attachment_name"`
	}
)
