package seeders

import (
	"kumande/factories"
	"kumande/modules/user"
	userWeather "kumande/modules/user_weather"
	"log"
)

func SeedUserWeathers(repo userWeather.UserWeatherRepository, userRepo user.UserRepository, count int) {
	// Empty Table
	repo.DeleteAll()

	// Fill Table
	var success = 0
	for i := 0; i < count; i++ {
		userWeather := factories.UserWeatherFactory()
		user, err := userRepo.FindOneRandom()
		if err != nil {
			log.Printf("failed to seed user weather %d: %v\n", i, err)
		}

		err = repo.CreateUserWeather(&userWeather, user.ID)
		if err != nil {
			log.Printf("failed to seed user weather %d: %v\n", i, err)
		}
		success++
	}
	log.Printf("Seeder : Success to seed %d User Weather", success)
}
