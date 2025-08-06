package seeders

import (
	"kumande/factories"
	"kumande/modules/consume"
	"kumande/modules/user"
	"log"
)

func SeedConsumeRates(repo consume.ConsumeRateRepository, userRepo user.UserRepository,
	consumeRepo consume.ConsumeRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	user, err := userRepo.FindOneHasConsumeAndConsumeList()
	if err != nil {
		log.Printf("failed to seed consume rate %v\n", err)
	}

	for _, dt := range user {
		for i := 0; i < count; i++ {
			consume, err := consumeRepo.FindOneRandom(dt.ID)
			if err != nil {
				log.Printf("failed to seed consume rate - %s at idx - %d : %v\n", dt.Username, i, err)
			}

			consumeRate := factories.ConsumeRateFactory(consume.ID)
			err = repo.CreateConsumeRate(&consumeRate, dt.ID)
			if err != nil {
				log.Printf("failed to seed consume rate - %s at idx - %d : %v\n", dt.Username, i, err)
			}
			success++
		}
	}

	log.Printf("Seeder : Success to seed %d consume rate", success)
}
