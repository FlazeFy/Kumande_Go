package userTrack

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User Track Interface
type UserTrackRepository interface {
	FindAllUserTrack(pagination utils.Pagination, userID uuid.UUID) ([]models.UserTrack, int64, error)
	CreateUserTrack(track *models.UserTrack, userID uuid.UUID) error

	// For Seeder
	DeleteAll() error
}

// User Track Struct
type userTrackRepository struct {
	db *gorm.DB
}

// User Track Constructor
func NewUserTrackRepository(db *gorm.DB) UserTrackRepository {
	return &userTrackRepository{db: db}
}

func (r *userTrackRepository) FindAllUserTrack(pagination utils.Pagination, userID uuid.UUID) ([]models.UserTrack, int64, error) {
	// Model
	var total int64
	var userTrack []models.UserTrack

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("user_tracks").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("user_tracks").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&userTrack)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(userTrack) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return userTrack, total, nil
}

func (r *userTrackRepository) CreateUserTrack(track *models.UserTrack, userID uuid.UUID) error {
	track.ID = uuid.New()
	track.CreatedAt = time.Now()
	track.CreatedBy = userID

	// Query
	return r.db.Create(track).Error
}

// For Seeder
func (r *userTrackRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.UserTrack{}).Error
}
