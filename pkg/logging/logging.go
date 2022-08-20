package logging

import (
	"golang-service-template/pkg/cfg"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	Log.SetLevel(logrus.DebugLevel)
	Log.SetFormatter(&logrus.TextFormatter{
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
			Log.Out = file
		} else {
			Log.Warning("Failed to log to file, using default stderr")
		}
	}
}
