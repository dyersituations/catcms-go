package setting

import (
	"catcms-go/database"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/labstack/echo/v4"

	model "catcms-go/model"
)

// Handles upserting an application setting to the datastore
func UpsertSetting(c echo.Context) error {
	// Get values from request body
	var key = c.FormValue("key")
	var value = c.FormValue("value")
	if key == "" || value == "" {
		return c.String(http.StatusBadRequest, "Missing setting key and/or value")
	}

	// Create setting object
	var setting model.Setting = model.Setting{
		AppId: os.Getenv("APP_ID"),
		Key:   key,
		Value: value,
	}

	// Find existing setting key if it exists
	// Existing datastore key needed for upsert
	var settings []model.Setting
	keys, err := database.Get(model.SettingKind, "Key", key, &settings)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error encountered")
	}
	var entityKey *datastore.Key = nil
	if len(keys) == 1 {
		entityKey = keys[0]
	}

	// Put the datastore entity
	err = database.Put(model.SettingKind, &setting, entityKey)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error encountered")
	}

	return c.String(http.StatusOK, "Successful upsert")
}

// Handles getting an application setting from the datastore
func GetSetting(c echo.Context) error {
	// Get the value from the request body
	var key = c.FormValue("key")
	if key == "" {
		return c.String(http.StatusBadRequest, "Missing setting key")
	}

	// Get the existing setting from the datastore
	var settings []model.Setting
	if keys, err := database.Get(model.SettingKind, "Key", key, &settings); err != nil {
		return c.String(http.StatusInternalServerError, "Error encountered")
	} else if len(keys) == 0 {
		return c.String(http.StatusBadRequest, "Setting not found")
	}

	return c.JSON(http.StatusOK, settings)
}
