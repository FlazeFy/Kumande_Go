package reminder

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderRepository interface {
	// For Seeder
	CreateReminder(reminder *models.Reminder, userId uuid.UUID) error
	DeleteAll() error
}

type reminderRepository struct {
	db *gorm.DB
}

func NewReminderRepository(db *gorm.DB) ReminderRepository {
	return &reminderRepository{db: db}
}

// For Seeder
func (r *reminderRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Reminder{}).Error
}
func (r *reminderRepository) CreateReminder(reminder *models.Reminder, userId uuid.UUID) error {
	reminder.ID = uuid.New()
	reminder.CreatedAt = time.Now()
	reminder.CreatedBy = userId

	// Query
	return r.db.Create(reminder).Error
}
