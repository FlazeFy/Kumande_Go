package seeders

import (
	"kumande/factories"
	"kumande/modules/nutrition"
	"kumande/modules/user"
	"log"
)

func SeedNutritions(repo nutrition.NutritionRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		nutrition := factories.NutritionFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed nutrition %d: %v\n", i, err)
		}

		err = repo.CreateNutrition(&nutrition, user.ID)
		if err != nil {
			log.Printf("failed to seed nutrition %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Nutrition", success)
}
