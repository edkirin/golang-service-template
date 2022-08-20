package main

import (
	"golang-service-template/pkg/cfg"
	"golang-service-template/pkg/controller"
	"golang-service-template/pkg/db"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func initLogger() {
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
		DisableQuote:    true,
	})

	if cfg.Config.Application.LogPath != nil {
		logPath := *cfg.Config.Application.LogPath
		file, err := os.OpenFile(
			filepath.Join(logPath, "service.log"),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666,
		)
		if err == nil {
			log.Out = file
		} else {
			log.Warning("Failed to log to file, using default stderr")
		}
	}
}

func main() {
	cfg.InitConfig(log)
	initLogger()
	db.InitDB(log)
	controller.Serve(log)
}
