package database

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

const projectID = "catcms"

// Stores an entity in the datastore
func Put(kind string, entities interface{}, key *datastore.Key) error {
	// Connect to the datastore
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, projectID)
	defer dsClient.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	// Create new key for an insert
	if key == nil {
		key = datastore.IncompleteKey(kind, nil)
	}

	// Put the entity
	if _, err := dsClient.Put(ctx, key, entities); err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

// Gets an entity from the datastore
func Get(appId string, kind string, filterKey string, filterValue string, entities interface{}) ([]*datastore.Key, error) {
	// Connect to the datastore
	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, projectID)
	defer dsClient.Close()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Create the query for the datastore
	// Always filter based on app ID
	q := datastore.NewQuery(kind).
		Filter("AppId =", appId).
		Filter(filterKey+" =", filterValue)

	// Get the entity
	if key, err := dsClient.GetAll(ctx, q, entities); err != nil {
		log.Println(err.Error())
		return nil, err
	} else {
		return key, nil
	}
}
