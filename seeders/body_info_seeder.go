package seeders

import (
	"kumande/factories"
	bodyInfo "kumande/modules/body_info"
	"kumande/modules/user"
	"log"
)

func SeedBodyInfo(repo bodyInfo.BodyInfoRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		bodyInfo := factories.BodyInfoFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed body info %d: %v\n", i, err)
		}

		err = repo.CreateBodyInfo(&bodyInfo, user.ID)
		if err != nil {
			log.Printf("failed to seed body info %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d body info", success)
}
