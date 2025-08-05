package factories

import (
	"kumande/models"

	"github.com/google/uuid"
)

func ReminderUsedFactory(reminderID uuid.UUID) models.ReminderUsed {
	return models.ReminderUsed{
		ReminderId: reminderID,
	}
}
