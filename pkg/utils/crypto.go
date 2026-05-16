package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"math/big"
)

const (
	OTPAlphaNum = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	OTPNumeric  = "0123456789"
	OTPHex      = "0123456789abcdef"
	OTPComplex  = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
)

// GenerateOTP generates a random OTP from the given charset.
func GenerateOTP(length int, charset string) string {
	result := make([]byte, length)
	for i := range result {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[n.Int64()]
	}
	return string(result)
}

// GenerateAlNumOTP generates an alphanumeric OTP.
func GenerateAlNumOTP(length int) string {
	return GenerateOTP(length, OTPAlphaNum)
}

// GenerateNumericOTP generates a numeric-only OTP.
func GenerateNumericOTP(length int) string {
	return GenerateOTP(length, OTPNumeric)
}

// GenerateOTPByType generates an OTP based on the specified type.
func GenerateOTPByType(length int, otpType string) string {
	switch otpType {
	case "num":
		return GenerateOTP(length, OTPNumeric)
	case "hex":
		return GenerateOTP(length, OTPHex)
	case "complex":
		return GenerateOTP(length, OTPComplex)
	case "alnum":
		return GenerateOTP(length, OTPAlphaNum)
	default:
		return GenerateOTP(length, OTPAlphaNum)
	}
}

// HashString returns the SHA-256 hex hash of a string.
func HashString(s string) string {
	h := sha256.Sum256([]byte(s))
	return hex.EncodeToString(h[:])
}
