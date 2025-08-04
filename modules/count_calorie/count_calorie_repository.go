package countCalorie

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CountCalorieRepository interface {
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
