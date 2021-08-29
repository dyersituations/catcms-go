package main

import (
	"log"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	auth "catcms-go/auth"
	service "catcms-go/service"
)

func main() {
	// Load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create Echo instance and add basic middleware
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Login route
	e.POST("/login", auth.Login)

	// Endpoint for getting a setting from the database
	e.GET("/get-setting", service.GetSetting)

	// Create group for auth
	g := e.Group("")

	// Configure middleware for JWT
	config := middleware.JWTConfig{
		Claims:     &jwt.StandardClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}
	g.Use(middleware.JWTWithConfig(config))

	// Endpoint for upserting a setting to the database
	g.POST("/upsert-setting", service.UpsertSetting)

	// Listen to port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
