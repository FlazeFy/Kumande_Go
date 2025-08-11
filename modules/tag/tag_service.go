package tag

import "github.com/google/uuid"

// Tag Interface
type TagService interface {
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

func (r *tagService) HardDeleteTagByID(ID, userID uuid.UUID) error {
	return r.tagRepo.HardDeleteTagByID(ID, userID)
}
