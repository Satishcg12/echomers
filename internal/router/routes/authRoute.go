package routes

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satishcg12/echomers/internal/handler"
	"github.com/satishcg12/echomers/internal/repos"
	"github.com/satishcg12/echomers/internal/services"
	"github.com/satishcg12/echomers/internal/utils"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(e *echo.Group, db *gorm.DB) {
	repo := repos.NewUserRepository(db)
	authService := services.NewAuthService(repo)
	// convert to int
	smtpPort, _ := strconv.Atoi(utils.GetEnv("SMTP_PORT"))
	mailService := services.NewEmailService(
		utils.GetEnv("SMTP_HOST"),
		smtpPort,
		utils.GetEnv("SMTP_USERNAME"),
		utils.GetEnv("SMTP_PASSWORD"),
	)

	handler := handler.NewAuthHandler(authService, mailService)

	auth := e.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)

		auth.PUT("/verify-email", handler.VerifyEmail)
		auth.PUT("/reset-password", handler.ResetPassword)
	}
}
