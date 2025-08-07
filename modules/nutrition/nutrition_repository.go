package nutrition

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Nutrition Interface
type NutritionRepository interface {
	CreateNutrition(nutrition *models.Nutrition, userID uuid.UUID) error
	DeleteAll() error
}

// Nutrition Struct
type nutritionRepository struct {
	db *gorm.DB
}

// Nutrition Constructor
func NewNutritionRepository(db *gorm.DB) NutritionRepository {
	return &nutritionRepository{db: db}
}

// For Seeder
func (r *nutritionRepository) CreateNutrition(nutrition *models.Nutrition, userID uuid.UUID) error {
	// Default
	nutrition.ID = uuid.New()
	nutrition.CreatedAt = time.Now()
	nutrition.CreatedBy = userID

	// Query
	return r.db.Create(nutrition).Error
}

func (r *nutritionRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Nutrition{}).Error
}
