package seeders

import (
	"kumande/factories"
	"kumande/modules/tag"
	"kumande/modules/user"
	"log"
)

func SeedTags(repo tag.TagRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		dct := factories.TagFactory()

		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed tag %v\n", err)
		}
		err = repo.CreateTag(&dct, user.ID)
		if err != nil {
			log.Printf("failed to seed tag %v\n", err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Tag", success)
}
