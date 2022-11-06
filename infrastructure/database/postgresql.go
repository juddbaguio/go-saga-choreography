package database

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/juddbaguio/go-saga-choreography/business/data/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type config struct {
	Host     string `env:"PSQL_HOST" envDefault:"localhost"`
	Username string `env:"PSQL_USERNAME,required"`
	Password string `env:"PSQL_PASSWORD,required"`
	Database string `env:"PSQL_DATABASE,required"`
	Port     int    `env:"PSQL_PORT" envDefault:"5432"`
}

func ConnectPsqlDB() (*gorm.DB, error) {
	cfg := &config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateBooking(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(&entities.Booking{})
}

func MigrateMovie(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(&entities.Movie{})
}

func MigratePayment(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(&entities.Payment{})
}

func MigrateCinema(dbConn *gorm.DB) error {
	return dbConn.AutoMigrate(
		&entities.Cinema{},
		&entities.CinemaMovie{},
		&entities.CinemaSeat{},
	)
}
