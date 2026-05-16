package routes

import (
	"embed"
	"io/fs"
	"net/http"
	"github.com/gin-gonic/gin"

	"/internal/handlers"
	"/internal/middleware"
	"/pkg/utils"
)

// SetupRouter creates and configures the Gin router.
func SetupRouter(h *handlers.WebHandlers, u *utils.Utils, webFS embed.FS) *gin.Engine {
	if !u.Config.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	// Static files from embedded FS
	staticFS, _ := fs.Sub(webFS, "web/template/default/static")
	r.StaticFS("/static", http.FS(staticFS))

	// Load HTML templates
	u.LoadTemplates(webFS)
	r.SetHTMLTemplate(u.Templates)

	// Public routes
	r.GET("/", h.HomeHandler)
	r.GET("/ping", h.PingHandler)
	r.GET("/login", h.LoginPageHandler)

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
		// Add protected routes here
	}

	return r
}
