package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamhanif11/shortlink-backend.git/internal/dto"
	"github.com/iamhanif11/shortlink-backend.git/internal/service"
	"github.com/iamhanif11/shortlink-backend.git/pkg"
)

type LinkController struct {
	linkService *service.LinkService
}

func NewLinkController(linkService *service.LinkService) *LinkController {
	return &LinkController{
		linkService: linkService,
	}
}

// @Summary      Create a new short link
// @Description  Authenticated users can transform a long destination URL into a clean, short link. If custom slug is omitted, the system will auto-generate a random unique identifier.
// @Tags         Link Management
// @Accept       json
// @Produce      json
// @Param        request  body      dto.CreateLinkReq  true  "Create Link Payload (destination URL and optional custom slug)"
// @Success      201      {object}  dto.Response[dto.LinkDetailRes] "Shortlink created successfully"
// @Failure      400      {object}  dto.ErrorResponse               "Bad Request - Validation error, reserved keyword used, or duplicate slug"
// @Failure      401      {object}  dto.ErrorResponse               "Unauthorized - Token missing or invalid signature"
// @Failure      500      {object}  dto.ErrorResponse               "Internal Server Error - Invalid token payload context format"
// @Router       /api/links [post]
// @Security     ApiKeyAuth
func (l *LinkController) CreateLink(ctx *gin.Context) {
	token, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "Unauthorized: Token not exist",
			Success: false,
		})
		return
	}

	claims, ok := token.(*pkg.Claims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Unauthorized: Format token Invalid",
			Success: false,
		})
		return
	}

	var body dto.CreateLinkReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Println("Error: ", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid Validation",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	res, err := l.linkService.CreateShortLink(ctx.Request.Context(), claims.Id, body)
	if err != nil {
		log.Println("error: ", err.Error())
		ctx.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Message: "Invalid Validation",
			Success: false,
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, dto.Response[dto.LinkDetailRes]{
		Message: "Shortlink created Succesfully",
		Success: true,
		Results: res,
	})
}

// @Summary      Get all short links for the authenticated user
// @Description  Retrieve a list of all active short links created by the logged-in user.
// @Tags         Link Management
// @Accept       json
// @Produce      json
// @Success      200      {object}  dto.Response[[]dto.LinkDetailRes] "List of user shortlinks retrieved successfully"
// @Failure      401      {object}  dto.ErrorResponse                 "Unauthorized - Token missing or invalid"
// @Failure      500      {object}  dto.ErrorResponse                 "Internal Server Error - Database or context error"
// @Router       /api/links [get]
// @Security     ApiKeyAuth
func (l *LinkController) GetUserLinks(ctx *gin.Context) {
	token, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, dto.ErrorResponse{
			Message: "Unauthorized: Token not Exist",
			Success: false,
		})
		return
	}

	claims, ok := token.(*pkg.Claims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Unauthorized: Format Token Invalid",
			Success: false,
		})
		return
	}

	res, err := l.linkService.GetUserLinks(ctx.Request.Context(), claims.Id)
	if err != nil {
		log.Println("Error fetching user links:", err.Error())
		ctx.JSON(http.StatusInternalServerError, dto.ErrorResponse{
			Message: "Internal server error failed to fetch links",
			Success: false,
		})
		return
	}
	ctx.JSON(http.StatusOK, dto.Response[[]dto.LinkDetailRes]{
		Message: "User links retrieved succesfully",
		Success: true,
		Results: res,
	})
}
