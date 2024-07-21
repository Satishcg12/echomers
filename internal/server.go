package internal

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/satishcg12/echomers/internal/database"
	"github.com/satishcg12/echomers/internal/router"
	"github.com/satishcg12/echomers/internal/utils"
	"github.com/satishcg12/echomers/internal/utils/validators"
)

type (
	ServerConfig struct {
		Host string
		Port string
	}
	Server struct {
		config ServerConfig
		e      *echo.Echo
	}

	ServerInterface interface {
		Start() error
	}
)

func NewServer(config ServerConfig) ServerInterface {
	e := echo.New()
	return &Server{
		config: config,
		e:      e,
	}
}

func (s *Server) Start() error {
	// init database
	db := database.NewDatabase(database.ConfigDatabase{
		Host:     utils.GetEnvOrDefault("DB_HOST", "localhost"),
		Port:     utils.GetEnvOrDefault("DB_PORT", "3306"),
		Username: utils.GetEnvOrDefault("DB_USERNAME", "root"),
		Password: utils.GetEnvOrDefault("DB_PASSWORD", "root"),
		Database: utils.GetEnvOrDefault("DB_NAME", "echomers"),
	})

	conn, err := db.Connect()
	if err != nil {
		return err
	}
	s.e.Use(middleware.Logger())
	s.e.Use(middleware.Recover())

	// init routes
	s.e.Validator = validators.NewValidator()
	router.Init(s.e, conn)

	// init server
	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	log.Printf("Server is running at %s", address)
	return s.e.Start(address)
}
