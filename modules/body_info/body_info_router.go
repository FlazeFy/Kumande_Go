package bodyInfo

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func BodyInfoRouter(r *gin.Engine, bodyInfoController BodyInfoController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_body_info_user := api.Group("/body_infos")
		protected_body_info_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_body_info_user.GET("/", bodyInfoController.GetAllBodyInfo)
			protected_body_info_user.DELETE("/:id", bodyInfoController.HardDeleteBodyInfoById, middlewares.AuditTrailMiddleware(db, "hard_delete_body_info_by_id"))
		}
	}
}
