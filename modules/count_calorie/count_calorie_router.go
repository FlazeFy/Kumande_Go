package countCalorie

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func CountCalorieRouter(r *gin.Engine, count_calorieController CountCalorieController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_count_calorie_user := api.Group("/count_calories")
		protected_count_calorie_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_count_calorie_user.GET("/last", count_calorieController.GetLastCountCalorie)
			protected_count_calorie_user.DELETE("/:id", count_calorieController.HardDeleteCountCalorieById, middlewares.AuditTrailMiddleware(db, "hard_delete_count_calorie_by_id"))
		}
	}
}
