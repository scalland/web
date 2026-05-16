package routes

import (
	"embed"
	"io/fs"
	"net/http"
	"github.com/gin-gonic/gin"

	"scalland/internal/handlers"
	"scalland/internal/middleware"
	"scalland/pkg/utils"
)

// SetupRouter creates and configures the Gin router.
func SetupRouter(h *handlers.WebHandlers, u *utils.Utils, webFS embed.FS) *gin.Engine {
	if !u.Config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.CORSMiddleware())

	// Static files from embedded FS
	staticFS, _ := fs.Sub(webFS, "web/template/scalland_20260518/static")
	r.StaticFS("/static", http.FS(staticFS))

	// Load HTML templates
	r.HTMLRender = u.LoadTemplates(webFS)

	// Public routes
	r.GET("/", h.HomeHandler)
	r.GET("/services.html", h.ServicesHandler)
	r.GET("/services", h.ServicesHandler)
	r.GET("/aboutus.html", h.AboutHandler)
	r.GET("/aboutus", h.AboutHandler)
	r.GET("/contactus.html", h.ContactHandler)
	r.GET("/contactus", h.ContactHandler)
	r.GET("/login.html", h.LoginPageHandler)
	r.GET("/login", h.LoginPageHandler)
	r.GET("/ping", h.PingHandler)

	// Auth API
	api := r.Group("/api/v1")
	{
		api.POST("/auth/send-otp", h.SendOTPHandler)
		api.POST("/auth/verify-otp", h.VerifyOTPHandler)
		api.POST("/auth/logout", h.LogoutHandler)
	}

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(u.Config.App.JWTSecret))
	{
		auth.GET("/dashboard", h.DashboardHandler)
		auth.GET("/dashboard.html", h.DashboardHandler)
	}

	return r
}
