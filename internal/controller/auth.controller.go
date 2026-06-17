package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// @Summary      Register a new user
// @Description  Create a new user account with email and password
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      dto.RegisterReq  true  "Registration Payload"
// @Success      201      {object}  dto.Response[dto.RegisterRes] "Registration successful"
// @Failure      400      {object}  dto.ErrorResponse             "Invalid input validation"
// @Failure      409      {object}  dto.ErrorResponse             "Email already registered"
// @Router       /api/register [post]
func (ac *AuthController) Register(ctx *gin.Context) {
	var body dto.RegisterReq
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid input validation",
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	res, err := ac.authService.RegisterUser(ctx.Request.Context(), body)
	if err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusConflict, dto.ErrorResponse{
			Message: "Email is Registered",
			Success: false,
			Error:   "Internal Server Error",
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response[dto.RegisterRes]{
		Results: res,
		Message: "Register Success",
		Success: true,
	})
}
