package seeders

import (
	"kumande/factories"
	"kumande/modules/budget"
	"kumande/modules/user"
	"log"
)

func SeedBudget(repo budget.BudgetRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		budget := factories.BudgetFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed budget %d: %v\n", i, err)
		}

		err = repo.CreateBudget(&budget, user.ID)
		if err != nil {
			log.Printf("failed to seed budget %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Budget", success)
}
