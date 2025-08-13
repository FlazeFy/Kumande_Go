package countCalorie

import (
	"kumande/models"

	"github.com/google/uuid"
)

// CountCalorie Interface
type CountCalorieService interface {
	GetLastCountCalorie(userID uuid.UUID) (*models.CountCalorie, error)
	HardDeleteCountCalorieByID(ID, userID uuid.UUID) error
}

// CountCalorie Struct
type countCalorieService struct {
	countCalorieRepo CountCalorieRepository
}

// CountCalorie Constructor
func NewCountCalorieService(countCalorieRepo CountCalorieRepository) CountCalorieService {
	return &countCalorieService{
		countCalorieRepo: countCalorieRepo,
	}
}

func (r *countCalorieService) GetLastCountCalorie(userID uuid.UUID) (*models.CountCalorie, error) {
	return r.countCalorieRepo.FindLastCountCalorie(userID)
}

func (r *countCalorieService) HardDeleteCountCalorieByID(ID, userID uuid.UUID) error {
	return r.countCalorieRepo.HardDeleteCountCalorieByID(ID, userID)
}
