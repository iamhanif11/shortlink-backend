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
// @Failure      400      {object}  dto.ErrorResponse             "Invalid input validation/ password mismatch"
// @Failure      409      {object}  dto.ErrorResponse             "Email already registered"
// @Router       /api/register [post]
func (ac *AuthController) Register(ctx *gin.Context) {
	var body dto.RegisterReq
	if err := ctx.ShouldBindWith(&body, binding.JSON); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Password and Confirm Password must match, and email must be valid",
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

// @Summary      User Login
// @Description  Authenticate user with email and password to get a JWT token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginReq  true  "Login Payload"
// @Success      200      {object}  dto.Response[dto.LoginResponse] "Login successful"
// @Failure      400      {object}  dto.ErrorResponse                 "Invalid input validation"
// @Failure      401      {object}  dto.ErrorResponse                 "Unauthorized - Invalid email or password"
// @Router       /api/login [post]
func (ac *AuthController) Login(ctx *gin.Context) {
	var body dto.LoginReq
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid input validation",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	res, err := ac.authService.LoginUser(ctx.Request.Context(), body)
	if err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "Login Failed",
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response[dto.LoginResponse]{
		Message: "Login Succesfully",
		Success: true,
		Results: res,
	})
}
