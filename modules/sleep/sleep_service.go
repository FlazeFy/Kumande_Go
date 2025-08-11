package sleep

import (
	"github.com/google/uuid"
)

// Sleep Interface
type SleepService interface {
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

func (r *sleepService) HardDeleteSleepByID(ID, userID uuid.UUID) error {
	return r.sleepRepo.HardDeleteSleepByID(ID, userID)
}
