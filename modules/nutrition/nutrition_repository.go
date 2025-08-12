package nutrition

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Nutrition Interface
type NutritionRepository interface {
	FindAllNutrition(pagination utils.Pagination, userID uuid.UUID) ([]models.Nutrition, int64, error)
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

func (r *nutritionRepository) FindAllNutrition(pagination utils.Pagination, userID uuid.UUID) ([]models.Nutrition, int64, error) {
	// Model
	var total int64
	var nutrition []models.Nutrition

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("nutritions").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("nutritions").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&nutrition)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(nutrition) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return nutrition, total, nil
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
