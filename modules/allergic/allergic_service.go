package allergic

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Allergic Interface
type AllergicService interface {
	GetAllAllergic(pagination utils.Pagination, userID uuid.UUID) ([]models.Allergic, int64, error)
	HardDeleteAllergicByID(ID, userID uuid.UUID) error
}

// Allergic Struct
type allergicService struct {
	allergicRepo AllergicRepository
}

// Allergic Constructor
func NewAllergicService(allergicRepo AllergicRepository) AllergicService {
	return &allergicService{
		allergicRepo: allergicRepo,
	}
}

func (s *allergicService) GetAllAllergic(pagination utils.Pagination, userID uuid.UUID) ([]models.Allergic, int64, error) {
	return s.allergicRepo.FindAllAllergic(pagination, userID)
}

func (r *allergicService) HardDeleteAllergicByID(ID, userID uuid.UUID) error {
	return r.allergicRepo.HardDeleteAllergicByID(ID, userID)
}
