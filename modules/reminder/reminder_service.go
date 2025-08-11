package reminder

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Reminder Interface
type ReminderService interface {
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
