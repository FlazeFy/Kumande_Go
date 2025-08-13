package reminder

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ReminderRouter(r *gin.Engine, reminderController ReminderController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_reminder_user := api.Group("/reminders")
		protected_reminder_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_reminder_user.GET("/most_context/:target_col", reminderController.GetMostContextReminder)
			protected_reminder_user.DELETE("/;id", reminderController.HardDeleteReminderByID)
		}

		// Private Routes - All Role
		protected_reminder_all := api.Group("/reminders")
		protected_reminder_all.Use(middlewares.AuthMiddleware(redisClient, "admin", "user"))
		{
			protected_reminder_all.GET("/", reminderController.GetAllReminder)
		}
	}
}
