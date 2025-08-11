package reminder

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderRepository interface {
	HardDeleteReminderByID(ID, userID uuid.UUID) error

	// For Seeder
	CreateReminder(reminder *models.Reminder, userId uuid.UUID) error
	DeleteAll() error
	FindOneRandom(userID uuid.UUID) (*models.Reminder, error)
}

type reminderRepository struct {
	db *gorm.DB
}

func NewReminderRepository(db *gorm.DB) ReminderRepository {
	return &reminderRepository{db: db}
}

func (r *reminderRepository) HardDeleteReminderByID(ID, userID uuid.UUID) error {
	// Query
	result := r.db.Unscoped().Where("id = ?", ID).Where("created_by = ?", userID).Delete(&models.Reminder{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
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
func (r *reminderRepository) FindOneRandom(userID uuid.UUID) (*models.Reminder, error) {
	var reminder models.Reminder

	err := r.db.Where("created_by", userID).Order("RANDOM()").Limit(1).First(&reminder).Error
	if err != nil {
		return nil, err
	}

	return &reminder, nil
}
