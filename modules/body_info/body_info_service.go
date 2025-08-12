package bodyInfo

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Body Info Interface
type BodyInfoService interface {
	GetAllBodyInfo(pagination utils.Pagination, userID uuid.UUID) ([]models.BodyInfo, int64, error)
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

func (s *bodyInfoService) GetAllBodyInfo(pagination utils.Pagination, userID uuid.UUID) ([]models.BodyInfo, int64, error) {
	return s.bodyInfoRepo.FindAllBodyInfo(pagination, userID)
}
