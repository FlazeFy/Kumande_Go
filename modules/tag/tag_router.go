package tag

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func TagRouter(r *gin.Engine, tagController TagController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_tag_user := api.Group("/tags")
		protected_tag_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_tag_user.GET("/", tagController.GetAllTag)
			protected_tag_user.DELETE("/:id", tagController.HardDeleteTagById, middlewares.AuditTrailMiddleware(db, "hard_delete_tag_by_id"))
		}
	}
}
