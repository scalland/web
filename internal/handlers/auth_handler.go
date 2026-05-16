package handlers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"scalland/pkg/utils"
)

// LoginPageHandler serves the login page.
func (h *WebHandlers) LoginPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.gohtml", gin.H{
		"Title":       "Login",
		"CurrentPage": "login",
	})
}

// SendOTPHandler sends an OTP to the user's email.
func (h *WebHandlers) SendOTPHandler(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	if !utils.IsValidEmail(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email"})
		return
	}

	u := h.Utils
	otp := utils.GenerateOTPByType(u.Config.OTP.Length, u.Config.OTP.OTPType)

	// Store OTP hash in DB
	otpHash := utils.HashString(otp)
	expiresAt := time.Now().Add(time.Duration(u.Config.OTP.ValidTill) * time.Second)

	_, err := u.DBWriter.Exec(
		"INSERT INTO otp_requests (email, otp_hash, expires_at) VALUES (?, ?, ?)",
		req.Email, otpHash, expiresAt,
	)
	if err != nil {
		u.Logger.Errorf("Failed to store OTP: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send OTP"})
		return
	}

	// Send OTP via email worker
	otpEmailTmpl := `<!DOCTYPE html>
<html><body style="font-family:sans-serif;background:#f4f6fb;padding:40px 0;">
<div style="max-width:480px;margin:0 auto;background:#fff;border-radius:12px;padding:40px;box-shadow:0 4px 24px rgba(10,25,47,0.08);">
  <h2 style="color:#0A192F;margin-bottom:8px;">Scalland Enterprise</h2>
  <p style="color:#555;margin-bottom:24px;">Use the following OTP to log in to your account. It is valid for 2 minutes.</p>
  <div style="font-size:36px;font-weight:700;letter-spacing:12px;color:#007BFF;background:#f0f6ff;border-radius:8px;padding:20px;text-align:center;margin-bottom:24px;">{{.OTP}}</div>
  <p style="color:#888;font-size:13px;">If you did not request this, please ignore this email.</p>
</div>
</body></html>`
	h.Worker.SendEmailHTMLAsync(req.Email, "Scalland Enterprise Login OTP", otpEmailTmpl, map[string]interface{}{
		"OTP":     otp,
		"AppName": "",
	})

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}

// VerifyOTPHandler verifies OTP and issues JWT tokens.
func (h *WebHandlers) VerifyOTPHandler(c *gin.Context) {
	var req struct {
		Email string `json:"email" binding:"required"`
		OTP   string `json:"otp" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and otp are required"})
		return
	}

	u := h.Utils
	otpHash := utils.HashString(req.OTP)

	// Verify OTP from DB
	var id int
	err := u.DBWriter.QueryRow(
		"SELECT id FROM otp_requests WHERE email = ? AND otp_hash = ? AND used = 0 AND expires_at > ? ORDER BY created_at DESC LIMIT 1",
		req.Email, otpHash, time.Now(),
	).Scan(&id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired OTP"})
		return
	}

	// Mark OTP as used
	u.DBWriter.Exec("UPDATE otp_requests SET used = 1 WHERE id = ?", id)

	// Upsert user
	u.DBWriter.Exec("INSERT INTO users (email) VALUES (?) ON DUPLICATE KEY UPDATE updated_at = NOW()", req.Email)

	// Get user ID
	var userID int
	u.DBWriter.QueryRow("SELECT id FROM users WHERE email = ?", req.Email).Scan(&userID)

	// Generate JWT token pair
	tokenPair, err := utils.GenerateTokenPair(userID, req.Email, "user", u.Config.App.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate tokens"})
		return
	}

	// Set cookies
	c.SetCookie("access_token", tokenPair.AccessToken, 3600, "/", "", false, true)
	c.SetCookie("refresh_token", tokenPair.RefreshToken, 604800, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message":       "login successful",
		"access_token":  tokenPair.AccessToken,
		"refresh_token": tokenPair.RefreshToken,
	})
}

// LogoutHandler clears auth cookies.
func (h *WebHandlers) LogoutHandler(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
