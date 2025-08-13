package tag

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository interface {
	FindAllTag(pagination utils.Pagination, userID uuid.UUID) ([]models.Tag, int64, error)
	HardDeleteTagByID(ID, userID uuid.UUID) error

	// For Seeder
	CreateTag(tag *models.Tag, userId *uuid.UUID) error
	DeleteAll() error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) FindAllTag(pagination utils.Pagination, userID uuid.UUID) ([]models.Tag, int64, error) {
	// Model
	var total int64
	var tag []models.Tag

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("tags").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("tags").
		Where("created_by = ?", userID).
		Or("created_by is null").
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&tag)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(tag) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return tag, total, nil
}

func (r *tagRepository) HardDeleteTagByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Tag{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// For Seeder
func (r *tagRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Tag{}).Error
}
func (r *tagRepository) CreateTag(tag *models.Tag, userId *uuid.UUID) error {
	tag.ID = uuid.New()
	tag.CreatedAt = time.Now()
	if userId != nil {
		tag.CreatedBy = userId
	}
	tag.TagSlug = utils.ConvertToSlug(tag.TagName)

	// Query
	return r.db.Create(tag).Error
}
