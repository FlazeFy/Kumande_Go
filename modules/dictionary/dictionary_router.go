package dictionary

import (
	"kumande/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func DictionaryRouter(r *gin.Engine, dictionaryController DictionaryController, redisClient *redis.Client, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		// Private Routes - All Roles
		protected_dictionary_all := api.Group("/dictionaries")
		protected_dictionary_all.Use(middlewares.AuthMiddleware(redisClient, "user", "admin"))
		{
			protected_dictionary_all.GET("/", dictionaryController.GetAllDictionary)
			protected_dictionary_all.GET("/:dictionary_type", dictionaryController.GetDictionaryByType)
		}

		// Private Routes - Admin
		protected_dictionary_admin := api.Group("/dictionaries")
		protected_dictionary_admin.Use(middlewares.AuthMiddleware(redisClient, "admin"))
		{
			protected_dictionary_admin.POST("/", dictionaryController.PostCreateDictionary, middlewares.AuditTrailMiddleware(db, "post_create_dictionary"))
			protected_dictionary_admin.DELETE("/:id", dictionaryController.HardDeleteDictionaryById, middlewares.AuditTrailMiddleware(db, "hard_delete_dictionary_by_id"))
		}
	}
}
