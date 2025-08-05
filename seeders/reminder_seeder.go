package seeders

import (
	"kumande/factories"
	"kumande/modules/reminder"
	"kumande/modules/user"
	"log"
)

func SeedReminder(repo reminder.ReminderRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		reminder := factories.ReminderFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed reminder %d: %v\n", i, err)
		}

		err = repo.CreateReminder(&reminder, user.ID)
		if err != nil {
			log.Printf("failed to seed reminder %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Reminder", success)
}
