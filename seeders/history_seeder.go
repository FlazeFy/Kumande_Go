package seeders

import (
	"kumande/factories"
	"kumande/modules/history"
	"kumande/modules/user"
	"log"
)

func SeedHistories(repo history.HistoryRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		history := factories.HistoryFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed history %d: %v\n", i, err)
		}

		err = repo.CreateHistory(&history, user.ID)
		if err != nil {
			log.Printf("failed to seed history %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d History", success)
}
