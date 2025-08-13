package seeders

import (
	"kumande/factories"
	"kumande/modules/tag"
	"kumande/modules/user"
	"log"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
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

		var userID *uuid.UUID
		if gofakeit.Bool() {
			userID = &user.ID
		} else {
			userID = nil
		}

		err = repo.CreateTag(&dct, userID)
		if err != nil {
			log.Printf("failed to seed tag %v\n", err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d Tag", success)
}
