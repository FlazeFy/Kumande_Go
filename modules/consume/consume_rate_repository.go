package consume

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ConsumeRate Interface
type ConsumeRateRepository interface {
	CreateConsumeRate(consumeRate *models.ConsumeRate, userID uuid.UUID) error
	DeleteAll() error
}

// ConsumeRate Struct
type consumeRateRepository struct {
	db *gorm.DB
}

// ConsumeRate Constructor
func NewConsumeRateRepository(db *gorm.DB) ConsumeRateRepository {
	return &consumeRateRepository{db: db}
}

// For Seeder
func (r *consumeRateRepository) CreateConsumeRate(consumeRate *models.ConsumeRate, userID uuid.UUID) error {
	// Default
	consumeRate.ID = uuid.New()
	consumeRate.CreatedAt = time.Now()
	consumeRate.CreatedBy = userID

	// Query
	return r.db.Create(consumeRate).Error
}

func (r *consumeRateRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.ConsumeRate{}).Error
}
