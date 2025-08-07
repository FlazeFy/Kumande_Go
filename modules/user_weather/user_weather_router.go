package userWeather

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserWeatherRouter(r *gin.Engine, userWeatherController UserWeatherController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_user_weathers_all := api.Group("/user_weathers")
		protected_user_weathers_all.Use(middlewares.AuthMiddleware(redisClient, "user"))
		{
			protected_user_weathers_all.GET("/most_context/:target_col", userWeatherController.GetMostContextUserWeather)
		}
	}
}
