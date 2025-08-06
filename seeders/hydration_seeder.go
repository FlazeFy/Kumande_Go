package seeders

import (
	"kumande/factories"
	"kumande/modules/hydration"
	"kumande/modules/user"
	"log"
)

func SeedHydrations(repo hydration.HydrationRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		hydration := factories.HydrationFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed hydration %d: %v\n", i, err)
		}

		err = repo.CreateHydration(&hydration, user.ID)
		if err != nil {
			log.Printf("failed to seed hydration %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Hydration", success)
}
