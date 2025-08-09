package allergic

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AllergicRouter(r *gin.Engine, allergicController AllergicController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_allergic_user := api.Group("/allergics")
		protected_allergic_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_allergic_user.DELETE("/:id", allergicController.HardDeleteAllergicById, middlewares.AuditTrailMiddleware(db, "hard_delete_allergic_by_id"))
		}
	}
}
