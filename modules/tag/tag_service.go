package tag

// Tag Interface
type TagService interface {
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
