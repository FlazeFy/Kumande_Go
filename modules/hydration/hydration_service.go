package hydration

import (
	"kumande/models"

	"github.com/google/uuid"
)

// Hydration Interface
type HydrationService interface {
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
