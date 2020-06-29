package datastore

import (
	"encoding/json"
	"io/ioutil"
	"log"

	datatype "tower/datatype/floorspec"
)

// FloorSpecStore ...
type FloorSpecStore struct {
	genericStore
}

// TODO: not necessary here since we gulp in all the floors

// InitFloorSpecStore ...
func InitFloorSpecStore() *FloorSpecStore {
	var store = new(FloorSpecStore)

	return store
}

// Read ...
func (door *FloorSpecStore) Read() (datatype.Floors, error) {
	var floors datatype.Floors

	// TODO: Temporarily loads a static floorspec file, eventually load from "database"
	jsonBytes, err := ioutil.ReadFile("../static/floorspec/floorspec.json")
	if err != nil {
		log.Print(err)
		return floors, err
	}

	if err := json.Unmarshal(jsonBytes, &floors); err != nil {
		log.Print(err)
		return floors, err
	}

	log.Printf("The floorspec is: %+v\n", floors)
	return floors, nil
}

// Write ...
func (door *FloorSpecStore) Write(floors *datatype.Floors) error {
	// TODO:
	return nil
}
