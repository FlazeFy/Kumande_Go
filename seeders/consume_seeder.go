package seeders

import (
	"kumande/factories"
	"kumande/modules/consume"
	"kumande/modules/user"
	"log"
)

func SeedConsume(repo consume.ConsumeRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		consume := factories.ConsumeFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed consume %d: %v\n", i, err)
		}

		err = repo.CreateConsume(&consume, user.ID)
		if err != nil {
			log.Printf("failed to seed consume %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d consume", success)
}
