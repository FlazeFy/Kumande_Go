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
	HardDeleteNutritionByID(ID, userID uuid.UUID) error
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

func (r *nutritionRepository) HardDeleteNutritionByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Nutrition{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
