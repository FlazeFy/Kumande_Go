package sleep

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Sleep Interface
type SleepRepository interface {
	FindAllSleep(pagination utils.Pagination, userID uuid.UUID) ([]models.Sleep, int64, error)
	HardDeleteSleepByID(ID, userID uuid.UUID) error

	// For Seeder
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

func (r *sleepRepository) FindAllSleep(pagination utils.Pagination, userID uuid.UUID) ([]models.Sleep, int64, error) {
	// Model
	var total int64
	var sleep []models.Sleep

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("sleeps").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("sleeps").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&sleep)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(sleep) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return sleep, total, nil
}

func (r *sleepRepository) HardDeleteSleepByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Sleep{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
