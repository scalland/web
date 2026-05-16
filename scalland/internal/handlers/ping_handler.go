package handlers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

var startTime = time.Now()

// PingHandler returns health check with uptime and version.
func (h *WebHandlers) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"uptime":  time.Since(startTime).String(),
		"version": "built-via-ldflags",
	})
}
