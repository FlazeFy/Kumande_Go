package seeders

import (
	"kumande/factories"
	"kumande/modules/allergic"
	"kumande/modules/user"
	"log"
)

func SeedAllergic(repo allergic.AllergicRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		allergic := factories.AllergicFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed allergic %d: %v\n", i, err)
		}

		err = repo.CreateAllergic(&allergic, user.ID)
		if err != nil {
			log.Printf("failed to seed allergic %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Allergic", success)
}
