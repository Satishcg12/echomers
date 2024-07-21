package database

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/satishcg12/echomers/internal/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	ConfigDatabase struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}
	Database struct {
		config ConfigDatabase
	}
	DatabaseInterface interface {
		Connect() (*gorm.DB, error)
		AutoMigrate(db *gorm.DB) error
	}
)

func NewDatabase(config ConfigDatabase) DatabaseInterface {
	return &Database{
		config: config,
	}
}

func (d *Database) Connect() (*gorm.DB, error) {
	// mysql connection
	dns := d.config.Username + ":" + d.config.Password + "@tcp(" + d.config.Host + ":" + d.config.Port + ")/" + d.config.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	// if flag -migration is set to true
	var migration bool
	flag.BoolVar(&migration, "migration", false, "Auto migrate database")
	flag.Parse()
	if migration {
		err = d.AutoMigrate(db)
		if err != nil {
			return nil, err
		}
		log.Println("Database migrated")
	}

	return db, nil
}

func (d *Database) AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&types.Users{},
	)
	if err != nil {
		return err
	}
	return nil

}
