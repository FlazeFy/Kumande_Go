package reminder

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reminder Interface
type ReminderService interface {
	GetAllReminder(pagination utils.Pagination, userID uuid.UUID) ([]models.Reminder, int64, error)
	HardDeleteReminderByID(ID, userID uuid.UUID) error
}

// Reminder Struct
type reminderService struct {
	reminderRepo     ReminderRepository
	reminderUsedRepo reminderUsedRepository
}

// Reminder Constructor
func NewReminderService(reminderRepo ReminderRepository) ReminderService {
	return &reminderService{
		reminderRepo: reminderRepo,
	}
}

func (s *reminderService) GetAllReminder(pagination utils.Pagination, userID uuid.UUID) ([]models.Reminder, int64, error) {
	return s.reminderRepo.FindAllReminder(pagination, userID)
}

func (r *reminderService) HardDeleteReminderByID(ID, userID uuid.UUID) error {
	// Service : Hard Delete Reminder By Reminder ID
	err := r.reminderRepo.HardDeleteReminderByID(ID, userID)
	if err != nil {
		return err
	}

	// Service : Hard Delete Reminder Used By Reminder Id
	err = r.reminderUsedRepo.HardDeleteReminderUsedByReminderID(ID, userID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	return nil
}
