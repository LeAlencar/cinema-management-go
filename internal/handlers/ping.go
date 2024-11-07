package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PingResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func Ping(c *gin.Context) {
	response := PingResponse{
		Message:   "pong",
		Timestamp: time.Now(),
	}
	c.JSON(http.StatusOK, response)
}
