package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("request has empty parameters")
	}

	// check if user already exists
	userExists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("there's an error checking if user exists: %w", err)
	}

	if userExists {
		return fmt.Errorf("a user with that username already exists")
	}

	// now user doesn't exist, insert into DB
	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("there's an error inserting user: %w", err)
	}
	return nil
}
