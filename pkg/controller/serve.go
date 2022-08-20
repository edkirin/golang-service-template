package controller

import (
	"golang-service-template/pkg/cfg"
	"golang-service-template/pkg/logging"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func corsMiddleware() gin.HandlerFunc {
	allowHeaders := [12]string{
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"accept",
		"origin",
		"Cache-Control",
		"X-Requested-With",
		"x-timezone",
		"Access-Control-Allow-Origin",
	}

	return func(c *gin.Context) {
		var allowOrigin string
		if cfg.Config.Application.IsProduction {
			allowOrigin = cfg.Config.Site.Url
		} else {
			allowOrigin = "*"
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(allowHeaders[:], ", "))
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func initRouter() *gin.Engine {
	if cfg.Config.Application.IsProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(corsMiddleware())
	router.Static("/data", "./data")

	router.GET("/ping", handlePing)
	return router
}

func Serve() {
	serverAddr := strings.Join([]string{cfg.Config.Server.Host, cfg.Config.Server.Port}, ":")
	router := initRouter()

	logging.Log.Info(strings.Join([]string{"Application started on ", serverAddr}, ""))
	router.Run(serverAddr)
}
