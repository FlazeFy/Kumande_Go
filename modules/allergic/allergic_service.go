package allergic

// Allergic Interface
type AllergicService interface {
}

// Allergic Struct
type allergicService struct {
	allergicRepo AllergicRepository
}

// Allergic Constructor
func NewAllergicService(allergicRepo AllergicRepository) AllergicService {
	return &allergicService{
		allergicRepo: allergicRepo,
	}
}
