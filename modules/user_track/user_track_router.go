package userTrack

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserTrackRouter(r *gin.Engine, userTrackController UserTrackController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - User
		protected_user_tracks_user := api.Group("/user_tracks")
		protected_user_tracks_user.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_user_tracks_user.GET("/most_context/:target_col", userTrackController.GetMostContextUserTrack)
		}

		// Private Routes - All Roles
		protected_user_tracks_all := api.Group("/user_tracks")
		protected_user_tracks_all.Use(middlewares.AuthMiddleware(redisClient, "admin", "user"))
		{
			protected_user_tracks_all.GET("/", userTrackController.GetAllUserTrack)
		}
	}
}
