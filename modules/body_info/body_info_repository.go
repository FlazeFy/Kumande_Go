package bodyInfo

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BodyInfoRepository interface {
	FindAllBodyInfo(pagination utils.Pagination, userID uuid.UUID) ([]models.BodyInfo, int64, error)
	HardDeleteBodyInfoByID(ID, userID uuid.UUID) error

	// For Seeder
	CreateBodyInfo(bodyInfo *models.BodyInfo, userId uuid.UUID) error
	DeleteAll() error
}

type bodyInfoRepository struct {
	db *gorm.DB
}

func NewBodyInfoRepository(db *gorm.DB) BodyInfoRepository {
	return &bodyInfoRepository{db: db}
}

func (r *bodyInfoRepository) FindAllBodyInfo(pagination utils.Pagination, userID uuid.UUID) ([]models.BodyInfo, int64, error) {
	// Model
	var total int64
	var bodyInfo []models.BodyInfo

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("body_infos").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("body_infos").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&bodyInfo)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(bodyInfo) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return bodyInfo, total, nil
}

func (r *bodyInfoRepository) HardDeleteBodyInfoByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.BodyInfo{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *bodyInfoRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.BodyInfo{}).Error
}
func (r *bodyInfoRepository) CreateBodyInfo(bodyInfo *models.BodyInfo, userId uuid.UUID) error {
	bodyInfo.ID = uuid.New()
	bodyInfo.CreatedAt = time.Now()
	bodyInfo.CreatedBy = userId

	// Query
	return r.db.Create(bodyInfo).Error
}
