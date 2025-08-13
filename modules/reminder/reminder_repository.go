package reminder

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReminderRepository interface {
	FindAllReminder(pagination utils.Pagination, userID uuid.UUID) ([]models.Reminder, int64, error)
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

func (r *reminderRepository) FindAllReminder(pagination utils.Pagination, userID uuid.UUID) ([]models.Reminder, int64, error) {
	// Model
	var total int64
	var reminder []models.Reminder

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("reminders").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("reminders").
		Where("created_by = ?", userID).
		Order("created_at DESC").
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&reminder)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(reminder) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return reminder, total, nil
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
