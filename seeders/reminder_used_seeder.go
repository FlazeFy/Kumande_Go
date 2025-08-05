package seeders

import (
	"kumande/factories"
	"kumande/modules/reminder"
	"kumande/modules/user"
	"log"
)

func SeedReminderUsed(repo reminder.ReminderUsedRepository, userRepo user.UserRepository, reminderRepo reminder.ReminderRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	user, err := userRepo.FindOneHasReminder()
	if err != nil {
		log.Printf("failed to seed reminder relation %v\n", err)
	}

	for _, dt := range user {
		for i := 0; i < count; i++ {
			reminder, err := reminderRepo.FindOneRandom(dt.ID)
			if err != nil {
				log.Printf("failed to seed reminder used relation user - %s at idx - %d : %v\n", dt.Username, i, err)
			}

			reminderUsed := factories.ReminderUsedFactory(reminder.ID)
			err = repo.CreateReminderUsed(&reminderUsed, dt.ID)
			if err != nil {
				log.Printf("failed to seed reminder used relation user - %s at idx - %d : %v\n", dt.Username, i, err)
			}
			success++
		}
	}

	log.Printf("Seeder : Success to seed %d reminder used relation", success)
}
