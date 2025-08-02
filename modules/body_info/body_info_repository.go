package bodyInfo

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BodyInfoRepository interface {
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
