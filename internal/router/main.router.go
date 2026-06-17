package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	//middleware global
	// router.Use(middleware.CORSMiddleware)
	//swagger docs
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// router.Static("/img/profile", "public/img/profiles")

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.ErrorResponse{
			Message: "Invalid Route",
			Success: false,
			Error:   "Not Found",
		})
	})
}
