package consume

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsumeListRepository interface {
	// For Seeder
	CreateConsumeList(consumeList *models.ConsumeList, userId uuid.UUID) error
	DeleteAll() error
}

type consumeListRepository struct {
	db *gorm.DB
}

func NewConsumeListRepository(db *gorm.DB) ConsumeListRepository {
	return &consumeListRepository{db: db}
}

// For Seeder
func (r *consumeListRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.ConsumeList{}).Error
}
func (r *consumeListRepository) CreateConsumeList(consumeList *models.ConsumeList, userId uuid.UUID) error {
	consumeList.ID = uuid.New()
	consumeList.CreatedAt = time.Now()
	consumeList.CreatedBy = userId
	consumeList.UpdatedAt = nil

	// Query
	return r.db.Create(consumeList).Error
}
