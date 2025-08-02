package budget

import (
	"kumande/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepository interface {
	// For Seeder
	CreateBudget(budget *models.Budget, userId uuid.UUID) error
	DeleteAll() error
}

type budgetRepository struct {
	db *gorm.DB
}

func NewBudgetRepository(db *gorm.DB) BudgetRepository {
	return &budgetRepository{db: db}
}

// For Seeder
func (r *budgetRepository) DeleteAll() error {
	return r.db.Where("1 = 1").Delete(&models.Budget{}).Error
}
func (r *budgetRepository) CreateBudget(budget *models.Budget, userId uuid.UUID) error {
	budget.ID = uuid.New()
	budget.CreatedAt = time.Now()
	budget.CreatedBy = userId
	budget.UpdatedAt = nil

	// Query
	return r.db.Create(budget).Error
}
