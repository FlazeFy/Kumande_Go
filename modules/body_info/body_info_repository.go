package bodyInfo

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BodyInfoRepository interface {
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
