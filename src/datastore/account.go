package datastore

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"tower/datatype"
)

// AccountStore ...
type AccountStore struct {
	genericStore
	AccountID string
}

// InitAccountStore ...
func InitAccountStore(accountID string) *AccountStore {
	var store = new(AccountStore)
	store.AccountID = accountID

	return store
}

// Read an account into the blob
// Uses the Blob.AccountID as unique identifier for acc
func (door *AccountStore) Read() (datatype.Account, error) {
	var acc datatype.Account

	// TODO: Temporarily loads a static user file, eventually load from "database"
	jsonBytes, err := ioutil.ReadFile("../static/user.json")
	if err != nil {
		log.Print(err)
		return acc, err
	}

	if err := json.Unmarshal(jsonBytes, &acc); err != nil {
		log.Print(err)
		return acc, err
	}

	log.Printf("The user is: %+v\n", acc)
	return acc, nil
}

// Write saves the blob account back to server
func (door *AccountStore) Write(acc *datatype.Account) error {
	// TODO: Write the data back to database/disk
	return nil
}
