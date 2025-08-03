package history

import (
	"kumande/models"

	"github.com/google/uuid"
)

// History Interface
type HistoryService interface {
	GetMyHistory(userID uuid.UUID) ([]models.History, error)
	HardDeleteHistoryByID(ID, createdBy uuid.UUID) error

	// For Task Scheduler
	DeleteHistoryForLastNDays(days int) (int64, error)
}

// History Struct
type historyService struct {
	historyRepo HistoryRepository
}

// History Constructor
func NewHistoryService(historyRepo HistoryRepository) HistoryService {
	return &historyService{
		historyRepo: historyRepo,
	}
}

func (r *historyService) GetMyHistory(userID uuid.UUID) ([]models.History, error) {
	return r.historyRepo.FindMyHistory(userID)
}

func (r *historyService) HardDeleteHistoryByID(ID, createdBy uuid.UUID) error {
	return r.historyRepo.HardDeleteHistoryByID(ID, createdBy)
}

// For Task Scheduler
func (r *historyService) DeleteHistoryForLastNDays(days int) (int64, error) {
	rows, err := r.historyRepo.DeleteHistoryForLastNDays(days)
	if err != nil {
		return 0, err
	}

	return rows, nil
}
