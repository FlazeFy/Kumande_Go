package allergic

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AllergicRepository interface {
	FindAllAllergic(pagination utils.Pagination, userID uuid.UUID) ([]models.Allergic, int64, error)

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

func (r *allergicRepository) FindAllAllergic(pagination utils.Pagination, userID uuid.UUID) ([]models.Allergic, int64, error) {
	// Model
	var total int64
	var allergic []models.Allergic

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("allergics").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("allergics").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&allergic)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(allergic) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return allergic, total, nil
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
