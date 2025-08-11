package bodyInfo

import "github.com/google/uuid"

// Body Info Interface
type BodyInfoService interface {
	HardDeleteBodyInfoByID(ID, userID uuid.UUID) error
}

// Body Info Struct
type bodyInfoService struct {
	bodyInfoRepo BodyInfoRepository
}

// Body Info Constructor
func NewBodyInfoService(bodyInfoRepo BodyInfoRepository) BodyInfoService {
	return &bodyInfoService{
		bodyInfoRepo: bodyInfoRepo,
	}
}

func (r *bodyInfoService) HardDeleteBodyInfoByID(ID, userID uuid.UUID) error {
	return r.bodyInfoRepo.HardDeleteBodyInfoByID(ID, userID)
}
