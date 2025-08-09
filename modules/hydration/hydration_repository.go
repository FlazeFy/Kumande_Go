package hydration

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hydration Interface
type HydrationRepository interface {
	CreateHydration(hydration *models.Hydration, userID uuid.UUID) error
	HardDeleteHydrationByID(ID, userID uuid.UUID) error
	DeleteAll() error
}

// Hydration Struct
type hydrationRepository struct {
	db *gorm.DB
}

// Hydration Constructor
func NewHydrationRepository(db *gorm.DB) HydrationRepository {
	return &hydrationRepository{db: db}
}

func (r *hydrationRepository) HardDeleteHydrationByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Hydration{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *hydrationRepository) CreateHydration(hydration *models.Hydration, userID uuid.UUID) error {
	// Default
	hydration.ID = uuid.New()
	hydration.CreatedAt = time.Now()
	hydration.CreatedBy = userID

	// Query
	return r.db.Create(hydration).Error
}

func (r *hydrationRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Hydration{}).Error
}
