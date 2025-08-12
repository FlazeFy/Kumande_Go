package nutrition

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Nutrition Interface
type NutritionService interface {
	GetAllNutrition(pagination utils.Pagination, userID uuid.UUID) ([]models.Nutrition, int64, error)
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

func (s *nutritionService) GetAllNutrition(pagination utils.Pagination, userID uuid.UUID) ([]models.Nutrition, int64, error) {
	return s.nutritionRepo.FindAllNutrition(pagination, userID)
}

func (r *nutritionService) HardDeleteNutritionByID(ID, userID uuid.UUID) error {
	return r.nutritionRepo.HardDeleteNutritionByID(ID, userID)
}
