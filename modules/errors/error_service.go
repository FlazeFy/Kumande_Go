package errors

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Error Interface
type ErrorService interface {
	GetAllError(pagination utils.Pagination) ([]models.ErrorAudit, int64, error)
	HardDeleteErrorByID(ID uuid.UUID) error

	// For Scheduler
	SchedulerGetAllErrorAudit() ([]models.ErrorAudit, error)
}

// Error Struct
type errorService struct {
	errorRepo ErrorRepository
}

// Error Constructor
func NewErrorService(errorRepo ErrorRepository) ErrorService {
	return &errorService{
		errorRepo: errorRepo,
	}
}

func (s *errorService) GetAllError(pagination utils.Pagination) ([]models.ErrorAudit, int64, error) {
	return s.errorRepo.FindAllError(pagination)
}

func (r *errorService) HardDeleteErrorByID(ID uuid.UUID) error {
	return r.errorRepo.HardDeleteErrorByID(ID)
}

// For Scheduler
func (s *errorService) SchedulerGetAllErrorAudit() ([]models.ErrorAudit, error) {
	return s.errorRepo.FindAllErrorAudit()
}
