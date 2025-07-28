package modules

import (
	"kumande/modules/admin"
	"kumande/modules/auth"
	"kumande/modules/user"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func SetUpDependency(r *gin.Engine, db *gorm.DB, redisClient *redis.Client) {
	// Dependency Repositories
	adminRepo := admin.NewAdminRepository(db)
	userRepo := user.NewUserRepository(db)

	// Dependency Services
	authService := auth.NewAuthService(userRepo, adminRepo, redisClient)

	// Dependency Controller
	authController := auth.NewAuthController(authService)

	// Routes Endpoint
	auth.AuthRouter(r, redisClient, *authController)
}
