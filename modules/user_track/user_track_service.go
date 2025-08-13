package userTrack

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// User Track Interface
type UserTrackService interface {
	GetAllUserTrack(pagination utils.Pagination, userID uuid.UUID) ([]models.UserTrack, int64, error)
}

// User Track Struct
type userTrackService struct {
	userTrackRepo UserTrackRepository
}

// User Track Constructor
func NewUserTrackService(userTrackRepo UserTrackRepository) UserTrackService {
	return &userTrackService{
		userTrackRepo: userTrackRepo,
	}
}

func (s *userTrackService) GetAllUserTrack(pagination utils.Pagination, userID uuid.UUID) ([]models.UserTrack, int64, error) {
	return s.userTrackRepo.FindAllUserTrack(pagination, userID)
}
