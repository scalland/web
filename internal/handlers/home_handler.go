package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HomeHandler serves the landing page.
func (h *WebHandlers) HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.gohtml", gin.H{
		"Title":       "",
		"AppName":     "",
		"CurrentPage": "home",
	})
}

// ServicesHandler serves the services page.
func (h *WebHandlers) ServicesHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "services.gohtml", gin.H{
		"Title":       "Services",
		"AppName":     "",
		"CurrentPage": "services",
	})
}

// AboutHandler serves the about us page.
func (h *WebHandlers) AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.gohtml", gin.H{
		"Title":       "About Us",
		"AppName":     "",
		"CurrentPage": "about",
	})
}

// ContactHandler serves the contact us page.
func (h *WebHandlers) ContactHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "contact.gohtml", gin.H{
		"Title":       "Contact Us",
		"AppName":     "",
		"CurrentPage": "contact",
	})
}

// DashboardHandler serves the protected dashboard page.
func (h *WebHandlers) DashboardHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
		"Title":       "Dashboard",
		"AppName":     "",
		"CurrentPage": "dashboard",
	})
}
