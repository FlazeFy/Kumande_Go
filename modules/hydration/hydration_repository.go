package hydration

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Hydration Interface
type HydrationRepository interface {
	FindAllHydration(pagination utils.Pagination, userID uuid.UUID) ([]models.Hydration, int64, error)
	FindHydrationByDate(userID uuid.UUID, date string) ([]models.Hydration, error)
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

func (r *hydrationRepository) FindAllHydration(pagination utils.Pagination, userID uuid.UUID) ([]models.Hydration, int64, error) {
	// Model
	var total int64
	var hydration []models.Hydration

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("hydrations").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("hydrations").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&hydration)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(hydration) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return hydration, total, nil
}

func (r *hydrationRepository) FindHydrationByDate(userID uuid.UUID, date string) ([]models.Hydration, error) {
	// Model
	var hydrations []models.Hydration

	// Query
	result := r.db.Where("created_by = ?", userID).
		Where("TO_CHAR(created_at, 'DD-MM-YYYY') = ?", date).
		Order("created_at DESC").
		Find(&hydrations)

	if len(hydrations) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	return hydrations, nil
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
