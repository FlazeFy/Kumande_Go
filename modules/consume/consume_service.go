package consume

// Consume Interface
type ConsumeService interface {
}

// Consume Struct
type consumeService struct {
	consumeRepo ConsumeRepository
}

// Consume Constructor
func NewConsumeService(consumeRepo ConsumeRepository) ConsumeService {
	return &consumeService{
		consumeRepo: consumeRepo,
	}
}
