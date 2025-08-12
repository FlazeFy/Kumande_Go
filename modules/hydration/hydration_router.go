package hydration

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func HydrationRouter(r *gin.Engine, hydrationController HydrationController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_hydration_user := api.Group("/hydrations")
		protected_hydration_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_hydration_user.GET("/by_date/:date", hydrationController.GetHydrationByDate)
			protected_hydration_user.POST("/", hydrationController.PostCreateHydration, middlewares.AuditTrailMiddleware(db, "post_create_hydration"))
			protected_hydration_user.DELETE("/:id", hydrationController.HardDeleteHydrationById, middlewares.AuditTrailMiddleware(db, "hard_delete_hydration_by_id"))
		}
	}
}
