package controller

import (
	"golang-service-template/pkg/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlePing(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		dto.PingDto{
			Message: "Pong!",
		},
	)
}
