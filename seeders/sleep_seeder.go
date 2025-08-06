package seeders

import (
	"kumande/factories"
	"kumande/modules/sleep"
	"kumande/modules/user"
	"log"
)

func SeedSleeps(repo sleep.SleepRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		sleep := factories.SleepFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed sleep %d: %v\n", i, err)
		}

		err = repo.CreateSleep(&sleep, user.ID)
		if err != nil {
			log.Printf("failed to seed sleep %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Sleep", success)
}
