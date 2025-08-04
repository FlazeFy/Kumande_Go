package seeders

import (
	"kumande/factories"
	countCalorie "kumande/modules/count_calorie"
	"kumande/modules/user"
	"kumande/utils"
	"log"
)

func SeedCountCalorie(repo countCalorie.CountCalorieRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		user, err := userRepo.FindOneRandom()
		age := utils.CalculateAge(user.BornAt)

		countCalorie := factories.CountCalorieFactory(age, user.Gender, user.ActivityFactor)
		if err != nil {
			log.Printf("failed to seed count calorie %d: %v\n", i, err)
		}

		err = repo.CreateCountCalorie(&countCalorie, user.ID)
		if err != nil {
			log.Printf("failed to seed count calorie %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d count calorie", success)
}
