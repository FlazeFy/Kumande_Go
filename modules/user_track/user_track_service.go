package userTrack

// User Track Interface
type UserTrackService interface {
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
