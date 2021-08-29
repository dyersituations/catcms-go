package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

// Logs in user and returns JWT token
func Login(c echo.Context) error {
	// Get username and password from request body
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Return unauthorized error
	if username != os.Getenv("BASIC_AUTH_USERNAME") || password != os.Getenv("BASIC_AUTH_PASSWORD") {
		return echo.ErrUnauthorized
	}

	// Set JWT claims with expiration
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	// Return JWT token
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
