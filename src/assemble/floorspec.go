package assemble

import (
	"log"
	"tower/datastore"
	datatype "tower/datatype/floorspec"
)

// TODO: get only floorspec for a specific floor

// FloorSpecAssembler ...
type FloorSpecAssembler struct {
	FloorID int
	Store   *datastore.FloorSpecStore
	Floors  *datatype.Floors
	Error   error
}

// InitFloorSpec ...
func InitFloorSpec(floorID int) *FloorSpecAssembler {
	var floorSpec = new(FloorSpecAssembler)
	floorSpec.FloorID = floorID
	floorSpec.Store = datastore.InitFloorSpecStore(floorID)

	return account
}

// AllFloors returns the floorspec for all floors
func (acc *AccountAssembler) AllFloors() error {
	// Read in the server side-saved account
	floors, err := acc.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	acc.Floors = &floors

	log.Print("successfully assembled floors")
	return nil
}
