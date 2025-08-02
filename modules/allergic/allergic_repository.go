package allergic

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AllergicRepository interface {
	// For Seeder
	CreateAllergic(allergic *models.Allergic, userId uuid.UUID) error
	DeleteAll() error
}

type allergicRepository struct {
	db *gorm.DB
}

func NewAllergicRepository(db *gorm.DB) AllergicRepository {
	return &allergicRepository{db: db}
}

// For Seeder
func (r *allergicRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Allergic{}).Error
}
func (r *allergicRepository) CreateAllergic(allergic *models.Allergic, userId uuid.UUID) error {
	allergic.ID = uuid.New()
	allergic.CreatedAt = time.Now()
	allergic.CreatedBy = userId
	allergic.UpdatedAt = nil

	// Query
	return r.db.Create(allergic).Error
}
