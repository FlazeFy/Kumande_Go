package consume

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConsumeRepository interface {
	// For Seeder
	CreateConsume(consume *models.Consume, userId uuid.UUID) error
	DeleteAll() error
	FindOneRandom(userID uuid.UUID) (*models.Consume, error)
}

type consumeRepository struct {
	db *gorm.DB
}

func NewConsumeRepository(db *gorm.DB) ConsumeRepository {
	return &consumeRepository{db: db}
}

// For Seeder
func (r *consumeRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Consume{}).Error
}
func (r *consumeRepository) CreateConsume(consume *models.Consume, userId uuid.UUID) error {
	consume.ID = uuid.New()
	consume.CreatedAt = time.Now()
	consume.CreatedBy = userId
	consume.UpdatedAt = nil
	consume.DeletedAt = nil

	// Query
	return r.db.Create(consume).Error
}
func (r *consumeRepository) FindOneRandom(userID uuid.UUID) (*models.Consume, error) {
	var consume models.Consume

	err := r.db.Where("created_by", userID).Order("RANDOM()").Limit(1).First(&consume).Error
	if err != nil {
		return nil, err
	}

	return &consume, nil
}
