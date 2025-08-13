package tag

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Tag Interface
type TagService interface {
	GetAllTag(pagination utils.Pagination, userID uuid.UUID) ([]models.Tag, int64, error)
	HardDeleteTagByID(ID, userID uuid.UUID) error
}

// Tag Struct
type tagService struct {
	tagRepo TagRepository
}

// Tag Constructor
func NewTagService(tagRepo TagRepository) TagService {
	return &tagService{
		tagRepo: tagRepo,
	}
}

func (s *tagService) GetAllTag(pagination utils.Pagination, userID uuid.UUID) ([]models.Tag, int64, error) {
	return s.tagRepo.FindAllTag(pagination, userID)
}

func (r *tagService) HardDeleteTagByID(ID, userID uuid.UUID) error {
	return r.tagRepo.HardDeleteTagByID(ID, userID)
}
