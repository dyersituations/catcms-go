package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	service "catcms-go/service"
)

func main() {
	// Load the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	// Endpoint for upserting a setting to the database
	e.POST("/upsert-setting", service.UpsertSetting)

	// Endpoint for getting a setting from the database
	e.GET("/get-setting", service.GetSetting)

	// Listen to port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
