package budget

// Budget Interface
type BudgetService interface {
}

// Budget Struct
type budgetService struct {
	budgetRepo BudgetRepository
}

// Budget Constructor
func NewBudgetService(budgetRepo BudgetRepository) BudgetService {
	return &budgetService{
		budgetRepo: budgetRepo,
	}
}
