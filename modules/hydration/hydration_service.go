package hydration

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Hydration Interface
type HydrationService interface {
	GetAllHydration(pagination utils.Pagination, userID uuid.UUID) ([]models.Hydration, int64, error)
	GetHydrationByDate(userID uuid.UUID, date string) ([]models.Hydration, error)
	CreateHydration(req models.Hydration, userID uuid.UUID) error
	HardDeleteHydrationByID(ID, userID uuid.UUID) error
}

// Hydration Struct
type hydrationService struct {
	hydrationRepo HydrationRepository
}

// Hydration Constructor
func NewHydrationService(hydrationRepo HydrationRepository) HydrationService {
	return &hydrationService{
		hydrationRepo: hydrationRepo,
	}
}

func (r *hydrationService) CreateHydration(hydration models.Hydration, userID uuid.UUID) error {
	return r.hydrationRepo.CreateHydration(&hydration, userID)
}

func (r *hydrationService) HardDeleteHydrationByID(ID, userID uuid.UUID) error {
	return r.hydrationRepo.HardDeleteHydrationByID(ID, userID)
}

func (s *hydrationService) GetAllHydration(pagination utils.Pagination, userID uuid.UUID) ([]models.Hydration, int64, error) {
	return s.hydrationRepo.FindAllHydration(pagination, userID)
}

func (r *hydrationService) GetHydrationByDate(userID uuid.UUID, date string) ([]models.Hydration, error) {
	return r.hydrationRepo.FindHydrationByDate(userID, date)
}
