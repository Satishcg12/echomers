package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satishcg12/echomers/internal/services"
	"github.com/satishcg12/echomers/internal/utils"
)

type (
	AuthHandler interface {
		Register(c echo.Context) error
		Login(c echo.Context) error
		VerifyEmail(c echo.Context) error
		ResetPassword(c echo.Context) error
	}

	authHandler struct {
		userService services.AuthService
		mailService services.EmailService
	}

	RegisterRequest struct {
		FullName        string `json:"full_name" form:"full_name" validate:"required,min=3,max=50"`
		Email           string `json:"email" form:"email" validate:"required,email,min=5,max=50"`
		Password        string `json:"password" form:"password" validate:"required,min=6,max=50"`
		ConfirmPassword string `json:"confirm_password" form:"confirm_password" validate:"required,eqfield=Password,min=6,max=50"`
	}
)

func NewAuthHandler(userService services.AuthService, mailService *services.EmailService) AuthHandler {
	return &authHandler{
		userService: userService,
		mailService: *mailService,
	}
}

func (h *authHandler) Register(c echo.Context) error {
	req := new(RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponse(http.StatusBadRequest, err.Error()))
	}
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewResponseWithData(http.StatusBadRequest, "Validation error", err.(*echo.HTTPError).Message))
	}

	if err := h.userService.Register(req.FullName, req.Email, req.Password); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, utils.NewResponse(http.StatusOK, "User registered successfully"))
}

func (h *authHandler) Login(c echo.Context) error {
	return nil
}

func (h *authHandler) VerifyEmail(c echo.Context) error {
	return nil
}

func (h *authHandler) ResetPassword(c echo.Context) error {
	return nil
}
