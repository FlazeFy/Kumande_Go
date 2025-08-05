package consume

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsumeListRelRepository interface {
	// For Seeder
	CreateConsumeListRel(consumeList *models.ConsumeListRelation, userId uuid.UUID) error
	DeleteAll() error
}

type consumeListRelRepository struct {
	db *gorm.DB
}

func NewConsumeListRelRepository(db *gorm.DB) ConsumeListRelRepository {
	return &consumeListRelRepository{db: db}
}

// For Seeder
func (r *consumeListRelRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.ConsumeListRelation{}).Error
}
func (r *consumeListRelRepository) CreateConsumeListRel(consumeListRel *models.ConsumeListRelation, userId uuid.UUID) error {
	consumeListRel.ID = uuid.New()
	consumeListRel.CreatedAt = time.Now()
	consumeListRel.CreatedBy = userId

	// Query
	return r.db.Create(consumeListRel).Error
}
