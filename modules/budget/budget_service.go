package budget

import (
	"kumande/models"
	"kumande/utils"

	"github.com/google/uuid"
)

// Budget Interface
type BudgetService interface {
	GetBudgetByYear(year string, userID uuid.UUID) ([]models.Budget, error)
	GetAllBudget(pagination utils.Pagination, userID uuid.UUID) ([]models.Budget, int64, error)
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

func (r *budgetService) GetAllBudget(pagination utils.Pagination, userID uuid.UUID) ([]models.Budget, int64, error) {
	return r.budgetRepo.FindAllBudget(pagination, userID)
}

func (r *budgetService) GetBudgetByYear(year string, userID uuid.UUID) ([]models.Budget, error) {
	return r.budgetRepo.FindBudgetByYear(year, userID)
}
