package countCalorie

import (
	"github.com/google/uuid"
)

// CountCalorie Interface
type CountCalorieService interface {
	HardDeleteCountCalorieByID(ID, userID uuid.UUID) error
}

// CountCalorie Struct
type dictionaryService struct {
	dictionaryRepo CountCalorieRepository
}

// CountCalorie Constructor
func NewCountCalorieService(dictionaryRepo CountCalorieRepository) CountCalorieService {
	return &dictionaryService{
		dictionaryRepo: dictionaryRepo,
	}
}

func (r *dictionaryService) HardDeleteCountCalorieByID(ID, userID uuid.UUID) error {
	return r.dictionaryRepo.HardDeleteCountCalorieByID(ID, userID)
}
