package budget

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BudgetRouter(r *gin.Engine, budgetController BudgetController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_budget_all := api.Group("/budgets")
		protected_budget_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_budget_all.GET("/", budgetController.GetAllBudget)
			protected_budget_all.GET("/:year", budgetController.GetBudgetByYear)
		}
	}
}
