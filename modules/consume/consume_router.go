package consume

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func ConsumeRouter(r *gin.Engine, consumeController ConsumeController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_consume_all := api.Group("/consumes")
		protected_consume_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_consume_all.GET("/most_context/:target_col", consumeController.GetMostContextConsume)
		}
	}
}
