package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/controller"
	"github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func LinkRouter(apiRouter *gin.RouterGroup, db *pgxpool.Pool, rc *redis.Client) {
	apiRouter.Use(middleware.VerifyToken)

	linkRepository := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepository, rc)
	linkController := controller.NewLinkController(linkService)

	apiRouter.POST("/links", linkController.CreateLink)
	apiRouter.GET("/links", linkController.GetUserLinks)
	apiRouter.DELETE("/links/:id", linkController.DeleteLink)
}
