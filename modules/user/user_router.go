package user

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func UserRouter(r *gin.Engine, userController UserController, redisClient *redis.Client) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_user_all := api.Group("/users")
		protected_user_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_user_all.GET("/my", userController.GetMyProfile)
		}
	}
}
