package reminder

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderUsedRepository interface {
	// For Seeder
	CreateReminderUsed(reminder *models.ReminderUsed, userId uuid.UUID) error
	DeleteAll() error
}

type reminderUsedRepository struct {
	db *gorm.DB
}

func NewReminderUsedRepository(db *gorm.DB) ReminderUsedRepository {
	return &reminderUsedRepository{db: db}
}

// For Seeder
func (r *reminderUsedRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.ReminderUsed{}).Error
}
func (r *reminderUsedRepository) CreateReminderUsed(reminder *models.ReminderUsed, userId uuid.UUID) error {
	reminder.ID = uuid.New()
	reminder.CreatedAt = time.Now()
	reminder.CreatedBy = userId

	// Query
	return r.db.Create(reminder).Error
}
