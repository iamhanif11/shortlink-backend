package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/iamhanif11/shortlink-backend.git/internal/cache"
	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/pkg"
	"github.com/redis/go-redis/v9"
)

func VerifyToken(rc *redis.Client) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Unauthorized access, token missing",
			})
			return
		}

		splittedBearer := strings.Split(bearerToken, " ")
		if len(splittedBearer) != 2 || strings.ToLower(splittedBearer[0]) != "bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Invalid authorization format",
			})
			return
		}
		token := splittedBearer[1]

		if cache.IsTokenBlacklisted(ctx.Request.Context(), rc, token) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Success: false,
				Message: "Your session has ended. Please login again.",
			})
			return
		}

		claims, err := new(pkg.Claims).VerifyJWT(token)
		if err != nil {
			log.Println("JWT Verification Error: ", err.Error())

			if errors.Is(err, jwt.ErrTokenInvalidIssuer) || errors.Is(err, jwt.ErrTokenExpired) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
					Success: false,
					Message: "Session expired or invalid, please re-login",
				})
				return
			}
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Success: false,
				Message: "Internal server error during authentication",
			})
			return
		}

		ctx.Set("claims", claims)
		ctx.Set("raw_token", token)
		ctx.Next()
	}
}
