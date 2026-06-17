package router

import (
	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/controller"
	"github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/iamhanif11/shortlink-backend.git/internal/repository"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func LinkRouter(apiRouter *gin.RouterGroup, db *pgxpool.Pool) {
	apiRouter.Use(middleware.VerifyToken)
	linkRepository := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepository)
	linkController := controller.NewLinkController(linkService)

	apiRouter.POST("/links", linkController.CreateLink)
	apiRouter.GET("/links", linkController.GetUserLinks)
}
