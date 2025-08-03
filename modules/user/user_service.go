package user

import (
	"kumande/models"

	"github.com/google/uuid"
)

// User Interface
type UserService interface {
	GetMyProfile(userID uuid.UUID) (*models.MyProfile, error)

	// Task Scheduler
	SchedulerGetUserReadyFetchWeather() ([]models.UserReadyFetchWeather, error)
}

// User Struct
type userService struct {
	userRepo UserRepository
}

// User Constructor
func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (r *userService) GetMyProfile(userID uuid.UUID) (*models.MyProfile, error) {
	// Repo : Find User By User Id
	return r.userRepo.FindById(userID)
}

func (r *userService) SchedulerGetUserReadyFetchWeather() ([]models.UserReadyFetchWeather, error) {
	// Repo : Find User Ready Fetch Weather
	rows, err := r.userRepo.FindUserReadyFetchWeather()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
