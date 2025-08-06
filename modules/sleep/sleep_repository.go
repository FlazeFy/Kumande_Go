package sleep

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Sleep Interface
type SleepRepository interface {
	CreateSleep(sleep *models.Sleep, userID uuid.UUID) error
	DeleteAll() error
}

// Sleep Struct
type sleepRepository struct {
	db *gorm.DB
}

// Sleep Constructor
func NewSleepRepository(db *gorm.DB) SleepRepository {
	return &sleepRepository{db: db}
}

// For Seeder
func (r *sleepRepository) CreateSleep(sleep *models.Sleep, userID uuid.UUID) error {
	// Default
	sleep.ID = uuid.New()
	sleep.CreatedAt = time.Now()
	sleep.CreatedBy = userID

	// Query
	return r.db.Create(sleep).Error
}

func (r *sleepRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Sleep{}).Error
}
