package bodyInfo

// Body Info Interface
type BodyInfoService interface {
}

// Body Info Struct
type budgetService struct {
	budgetRepo BodyInfoRepository
}

// Body Info Constructor
func NewBodyInfoService(budgetRepo BodyInfoRepository) BodyInfoService {
	return &budgetService{
		budgetRepo: budgetRepo,
	}
}
