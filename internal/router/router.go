package router

import (
	"github.com/labstack/echo/v4"
	"github.com/satishcg12/echomers/internal/router/routes"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to Echomers")
	})

	// group routes
	api := e.Group("/api")
	{
		// auth routes
		routes.RegisterAuthRoutes(api, db)

	}

}
