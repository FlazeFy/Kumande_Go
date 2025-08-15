package budget

import (
	"kumande/models"
	"kumande/utils"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BudgetRepository interface {
	FindAllBudget(pagination utils.Pagination, userID uuid.UUID) ([]models.Budget, int64, error)
	FindBudgetByYear(year string, userID uuid.UUID) ([]models.Budget, error)

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

func (r *budgetRepository) FindAllBudget(pagination utils.Pagination, userID uuid.UUID) ([]models.Budget, int64, error) {
	// Model
	var total int64
	var budget []models.Budget

	// Pagination Count
	offset := (pagination.Page - 1) * pagination.Limit
	countQuery := r.db.Table("budgets").Where("created_by = ?", userID)

	if err := countQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Model
	query := r.db.Table("budgets").
		Where("created_by = ?", userID).
		Order("budget_year DESC").
		Order(`
			CASE budget_month
				WHEN 'Jan' THEN 1
				WHEN 'Feb' THEN 2
				WHEN 'Mar' THEN 3
				WHEN 'Apr' THEN 4
				WHEN 'May' THEN 5
				WHEN 'Jun' THEN 6
				WHEN 'Jul' THEN 7
				WHEN 'Aug' THEN 8
				WHEN 'Sep' THEN 9
				WHEN 'Oct' THEN 10
				WHEN 'Nov' THEN 11
				WHEN 'Dec' THEN 12
			END DESC
		`).
		Limit(pagination.Limit).
		Offset(offset)

	result := query.Find(&budget)

	if result.Error != nil {
		return nil, 0, result.Error
	}
	if len(budget) == 0 {
		return nil, 0, gorm.ErrRecordNotFound
	}

	return budget, total, nil
}

func (r *budgetRepository) FindBudgetByYear(year string, userID uuid.UUID) ([]models.Budget, error) {
	// Model
	var budget []models.Budget

	// Model
	query := r.db.Table("budgets").
		Where("created_by = ?", userID).
		Where("budget_year = ?", year).
		Order(`
			CASE budget_month
				WHEN 'Jan' THEN 1
				WHEN 'Feb' THEN 2
				WHEN 'Mar' THEN 3
				WHEN 'Apr' THEN 4
				WHEN 'May' THEN 5
				WHEN 'Jun' THEN 6
				WHEN 'Jul' THEN 7
				WHEN 'Aug' THEN 8
				WHEN 'Sep' THEN 9
				WHEN 'Oct' THEN 10
				WHEN 'Nov' THEN 11
				WHEN 'Dec' THEN 12
			END DESC
		`)

	result := query.Find(&budget)

	if result.Error != nil {
		return nil, result.Error
	}
	if len(budget) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return budget, nil
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
