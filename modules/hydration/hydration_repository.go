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
