package userWeather

import (
	"kumande/models"

	"github.com/google/uuid"
)

// User Weather Interface
type UserWeatherService interface {
	// Task Scheduler
	CreateUserWeather(weather *models.UserWeather, userID uuid.UUID) error
}

// User Weather Struct
type userWeatherService struct {
	userWeatherRepo UserWeatherRepository
}

// User Weather Constructor
func NewUserWeatherService(userWeatherRepo UserWeatherRepository) UserWeatherService {
	return &userWeatherService{
		userWeatherRepo: userWeatherRepo,
	}
}

func (s *userWeatherService) CreateUserWeather(weather *models.UserWeather, userID uuid.UUID) error {
	// Repo : Create User Weather
	if err := s.userWeatherRepo.CreateUserWeather(weather, userID); err != nil {
		return err
	}

	return nil
}
