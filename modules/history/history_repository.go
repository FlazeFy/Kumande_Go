package history

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// History Interface
type HistoryRepository interface {
	FindMyHistory(userID uuid.UUID) ([]models.History, error)
	HardDeleteHistoryByID(ID, createdBy uuid.UUID) error

	// Task Scheduler
	DeleteHistoryForLastNDays(days int) (int64, error)

	// For Seeder
	CreateHistory(history *models.History, userID uuid.UUID) error
	DeleteAll() error
}

// History Struct
type historyRepository struct {
	db *gorm.DB
}

// History Constructor
func NewHistoryRepository(db *gorm.DB) HistoryRepository {
	return &historyRepository{db: db}
}

func (r *historyRepository) FindMyHistory(userID uuid.UUID) ([]models.History, error) {
	// Model
	var histories []models.History

	// Query
	if err := r.db.Where("created_by", userID).Find(&histories).Error; err != nil {
		return nil, err
	}

	return histories, nil
}

func (r *historyRepository) HardDeleteHistoryByID(ID, createdBy uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", createdBy).Delete(&models.History{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *historyRepository) CreateHistory(history *models.History, userID uuid.UUID) error {
	// Default
	history.ID = uuid.New()
	history.CreatedAt = time.Now()
	history.CreatedBy = userID

	// Query
	return r.db.Create(history).Error
}

// Task Scheduler
func (r *historyRepository) DeleteHistoryForLastNDays(days int) (int64, error) {
	// Cutoff Days
	cutoff := time.Now().AddDate(0, 0, -days)

	// Query
	result := r.db.Unscoped().Where("created_at < ?", cutoff).Delete(&models.History{})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func (r *historyRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.History{}).Error
}
