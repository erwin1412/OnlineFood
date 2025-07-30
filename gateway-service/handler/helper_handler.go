package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Helper functions to extract user information from JWT token context
func GetUserIDFromContext(c echo.Context) (string, error) {
	userID := c.Get("user_id")
	if userID == nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "User ID not found in token")
	}

	userIDStr, ok := userID.(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Invalid user ID format")
	}

	return userIDStr, nil
}

func GetEmailFromContext(c echo.Context) (string, error) {
	email := c.Get("email")
	if email == nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Email not found in token")
	}

	emailStr, ok := email.(string)
	if !ok {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Invalid email format")
	}

	return emailStr, nil
}
