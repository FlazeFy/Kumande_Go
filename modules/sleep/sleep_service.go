package sleep

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Sleep Interface
type SleepService interface {
	GetAllSleep(pagination utils.Pagination, userID uuid.UUID) ([]models.Sleep, int64, error)
	HardDeleteSleepByID(ID, userID uuid.UUID) error
}

// Sleep Struct
type sleepService struct {
	sleepRepo SleepRepository
}

// Sleep Constructor
func NewSleepService(sleepRepo SleepRepository) SleepService {
	return &sleepService{
		sleepRepo: sleepRepo,
	}
}

func (s *sleepService) GetAllSleep(pagination utils.Pagination, userID uuid.UUID) ([]models.Sleep, int64, error) {
	return s.sleepRepo.FindAllSleep(pagination, userID)
}

func (r *sleepService) HardDeleteSleepByID(ID, userID uuid.UUID) error {
	return r.sleepRepo.HardDeleteSleepByID(ID, userID)
}
