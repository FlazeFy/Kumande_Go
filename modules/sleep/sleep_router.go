package sleep

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SleepRouter(r *gin.Engine, sleepController SleepController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_sleep_user := api.Group("/sleeps")
		protected_sleep_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_sleep_user.GET("/", sleepController.GetAllSleep)
			protected_sleep_user.DELETE("/:id", sleepController.HardDeleteSleepById, middlewares.AuditTrailMiddleware(db, "hard_delete_sleep_by_id"))
		}
	}
}
