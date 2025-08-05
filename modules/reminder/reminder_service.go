package reminder

// Reminder Interface
type ReminderService interface {
}

// Reminder Struct
type reminderService struct {
	reminderRepo ReminderRepository
}

// Reminder Constructor
func NewReminderService(reminderRepo ReminderRepository) ReminderService {
	return &reminderService{
		reminderRepo: reminderRepo,
	}
}
