package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/controller"

	// "github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthRouter(apiRouter *gin.RouterGroup, db *pgxpool.Pool) {

	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	// authMiddleware := middleware.NewAuthMiddleware(authRepository)

	apiRouter.POST("/register", authController.Register)

}
