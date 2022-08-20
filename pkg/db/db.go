package db

import (
	"golang-service-template/pkg/cfg"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var log *logrus.Logger

func InitDB(logger *logrus.Logger) *gorm.DB {
	log = logger

	var connectionString = strings.Join([]string{
		"postgres://",
		cfg.Config.Database.Username, ":",
		cfg.Config.Database.Password, "@",
		cfg.Config.Database.Host, ":",
		cfg.Config.Database.Port, "/",
		cfg.Config.Database.Name,
		"?sslmode=disable",
	}, "")
	log.Info("Using database connection string: " + connectionString)

	var logLevel = gormLogger.Silent
	if cfg.Config.Application.DebugSQL {
		logLevel = gormLogger.Info
	}

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: gormLogger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Error("Error connecting to database: " + err.Error())
		return nil
	}

	return db
}
