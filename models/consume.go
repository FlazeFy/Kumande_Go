package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type (
	Consume struct {
		ID             uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey"`
		ConsumeName    string         `json:"consume_name" gorm:"type:varchar(75);not null" binding:"required,max=75"`
		ConsumeDetail  *string        `json:"consume_detail" gorm:"type:varchar(500);null" binding:"omitempty,max=500"`
		ConsumeType    string         `json:"consume_type" gorm:"type:varchar(14);null" binding:"omitempty,max=14"`
		ConsumeFrom    string         `json:"consume_from" gorm:"type:varchar(36);not null" binding:"required,max=36"`
		ConsumePrice   *int           `json:"consume_price" gorm:"type:int;null" binding:"omitempty,min=1,max=99999999"`
		ConsumeBuyAt   *time.Time     `json:"consume_buy_at" gorm:"type:date;null" binding:"required"`
		ConsumeQty     int            `json:"consume_qty" gorm:"type:int;not null" binding:"required,min=1"`
		ConsumeImage   *string        `json:"consume_image" gorm:"type:varchar(1000);null" binding:"omitempty,max=1000"`
		ConsumeProvide *string        `json:"consume_provide" gorm:"type:varchar(36);null" binding:"omitempty,max=36"`
		ConsumeCal     *int           `json:"consume_calorie" gorm:"type:int;null" binding:"omitempty"`
		ConsumeTag     datatypes.JSON `json:"consume_tag" gorm:"type:json;null" binding:"omitempty"`
		IsFavorite     bool           `json:"is_favorite" gorm:"type:boolean;not null" binding:"required"`
		CreatedAt      time.Time      `json:"created_at" gorm:"type:timestamp;not null"`
		UpdatedAt      *time.Time     `json:"updated_at" gorm:"type:timestamp;null"`
		DeletedAt      *time.Time     `json:"deleted_at" gorm:"type:timestamp;null"`
		// FK - User
		CreatedBy uuid.UUID `json:"created_by" gorm:"not null"`
		User      User      `json:"-" gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
	ConsumeTag struct {
		SlugName string `json:"slug_name"`
		TagName  string `json:"tag_name"`
	}
)
