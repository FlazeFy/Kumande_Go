package seeders

import (
	"kumande/factories"
	"kumande/modules/user"
	userTrack "kumande/modules/user_track"
	"log"
)

func SeedUserTracks(repo userTrack.UserTrackRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		track := factories.UserTrackFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed user track %d: %v\n", i, err)
		}

		err = repo.CreateUserTrack(&track, user.ID)
		if err != nil {
			log.Printf("failed to seed user track %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d User Track", success)
}
