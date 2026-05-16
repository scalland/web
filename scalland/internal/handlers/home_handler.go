package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HomeHandler serves the landing page.
func (h *WebHandlers) HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", gin.H{
		"Title":   "",
		"AppName": "",
	})
}
