package reminder

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderUsedRepository interface {
	HardDeleteReminderUsedByReminderID(ID, userID uuid.UUID) error

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

func (r *reminderUsedRepository) HardDeleteReminderUsedByReminderID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("reminder_id = ?", ID).Where("created_by = ?", userID).Delete(&models.ReminderUsed{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
