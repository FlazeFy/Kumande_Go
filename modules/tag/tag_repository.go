package tag

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository interface {
	// For Seeder
	CreateTag(tag *models.Tag, userId uuid.UUID) error
	DeleteAll() error
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

// For Seeder
func (r *tagRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Tag{}).Error
}
func (r *tagRepository) CreateTag(tag *models.Tag, userId uuid.UUID) error {
	tag.ID = uuid.New()
	tag.CreatedAt = time.Now()
	tag.CreatedBy = userId
	tag.TagSlug = utils.ConvertToSlug(tag.TagName)

	// Query
	return r.db.Create(tag).Error
}
