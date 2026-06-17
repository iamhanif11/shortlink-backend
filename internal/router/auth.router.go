package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/controller"
	"github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/redis/go-redis/v9"

	// "github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(apiRouter *gin.RouterGroup, db *pgxpool.Pool, rc *redis.Client) {

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, rc)
	authController := controller.NewAuthController(authService)

	// authMiddleware := middleware.NewAuthMiddleware(authRepository)

	apiRouter.POST("/register", authController.Register)
	apiRouter.POST("/login", authController.Login)

	apiRouter.DELETE("/logout", middleware.VerifyToken(rc), authController.Logout)
}
