package allergic

import "github.com/google/uuid"

// Allergic Interface
type AllergicService interface {
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

func (r *allergicService) HardDeleteAllergicByID(ID, userID uuid.UUID) error {
	return r.allergicRepo.HardDeleteAllergicByID(ID, userID)
}
