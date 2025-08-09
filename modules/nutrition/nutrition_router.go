package nutrition

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NutritionRouter(r *gin.Engine, nutritionController NutritionController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_nutrition_user := api.Group("/nutritions")
		protected_nutrition_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_nutrition_user.DELETE("/:id", nutritionController.HardDeleteNutritionById, middlewares.AuditTrailMiddleware(db, "hard_delete_nutrition_by_id"))
		}
	}
}
