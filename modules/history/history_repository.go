package history

import (
	"kumande/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// History Interface
type HistoryRepository interface {
	FindMyHistory(userID uuid.UUID) ([]models.History, error)
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
