package middleware

// RateLimiter defines the rate limiting interface.
type RateLimiter interface {
	Allow(key string) bool
}
