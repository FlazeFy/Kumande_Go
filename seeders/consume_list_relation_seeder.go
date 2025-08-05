package seeders

import (
	"kumande/factories"
	"kumande/modules/consume"
	"kumande/modules/user"
	"log"
)

func SeedConsumeListRelations(repo consume.ConsumeListRelRepository, userRepo user.UserRepository,
	consumeRepo consume.ConsumeRepository, consumeListRepo consume.ConsumeListRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	user, err := userRepo.FindOneHasConsumeAndConsumeList()
	if err != nil {
		log.Printf("failed to seed consumeList relation %v\n", err)
	}

	for _, dt := range user {
		for i := 0; i < count; i++ {
			consume, err := consumeRepo.FindOneRandom(dt.ID)
			if err != nil {
				log.Printf("failed to seed consume list relation user - %s at idx - %d : %v\n", dt.Username, i, err)
			}

			consumeList, err := consumeListRepo.FindOneRandom(dt.ID)
			if err != nil {
				log.Printf("failed to seed consume list relation user - %s at idx - %d : %v\n", dt.Username, i, err)
			}

			consumeListRel := factories.ConsumeListRelationFactory(consumeList.ID, consume.ID)
			err = repo.CreateConsumeListRel(&consumeListRel, dt.ID)
			if err != nil {
				log.Printf("failed to seed consume list relation user - %s at idx - %d : %v\n", dt.Username, i, err)
			}
			success++
		}
	}

	log.Printf("Seeder : Success to seed %d consume list relation", success)
}
