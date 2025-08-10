package countCalorie

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CountCalorieRepository interface {
	HardDeleteCountCalorieByID(ID, userID uuid.UUID) error

	// For Seeder
	CreateCountCalorie(countCalorie *models.CountCalorie, userId uuid.UUID) error
	DeleteAll() error
}

type countCalorieRepository struct {
	db *gorm.DB
}

func NewCountCalorieRepository(db *gorm.DB) CountCalorieRepository {
	return &countCalorieRepository{db: db}
}

func (r *countCalorieRepository) HardDeleteCountCalorieByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.CountCalorie{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *countCalorieRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.CountCalorie{}).Error
}
func (r *countCalorieRepository) CreateCountCalorie(count *models.CountCalorie, userId uuid.UUID) error {
	count.ID = uuid.New()
	count.CreatedAt = time.Now()
	count.CreatedBy = userId

	// Query
	return r.db.Create(count).Error
}
