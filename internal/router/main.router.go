package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	_ "github.com/iamhanif11/shortlink-backend.git/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine, db *pgxpool.Pool, rc *redis.Client) {
	//middleware global
	router.Use(middleware.CORSMiddleware)
	//swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// router.Static("/img/profile", "public/img/profiles")

	routeApi := router.Group("/api")
	AuthRouter(routeApi, db)
	LinkRouter(routeApi, db, rc)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "Invalid Route",
			Success: false,
			Error:   "Not Found",
		})
	})
}
