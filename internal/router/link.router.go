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

func LinkRouter(apiRouter *gin.RouterGroup, mainEngine *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {

	linkRepository := repository.NewLinkRepository(db)
	linkService := service.NewLinkService(linkRepository, rc)
	linkController := controller.NewLinkController(linkService)

	mainEngine.GET("/:slug", linkController.Redirect)

	protected := apiRouter.Group("")
	protected.Use(middleware.VerifyToken(rc))
	{
		protected.POST("/links", linkController.CreateLink)
		protected.GET("/links", linkController.GetUserLinks)
		protected.DELETE("/links/:id", linkController.DeleteLink)
	}
}
