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
	HardDeleteAllergicByID(ID, userID uuid.UUID) error
	DeleteAll() error
}

type allergicRepository struct {
	db *gorm.DB
}

func NewAllergicRepository(db *gorm.DB) AllergicRepository {
	return &allergicRepository{db: db}
}

func (r *allergicRepository) HardDeleteAllergicByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Allergic{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
