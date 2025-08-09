package nutrition

import (
	"github.com/google/uuid"
)

// Nutrition Interface
type NutritionService interface {
	HardDeleteNutritionByID(ID, userID uuid.UUID) error
}

// Nutrition Struct
type nutritionService struct {
	nutritionRepo NutritionRepository
}

// Nutrition Constructor
func NewNutritionService(nutritionRepo NutritionRepository) NutritionService {
	return &nutritionService{
		nutritionRepo: nutritionRepo,
	}
}

func (r *nutritionService) HardDeleteNutritionByID(ID, userID uuid.UUID) error {
	return r.nutritionRepo.HardDeleteNutritionByID(ID, userID)
}
