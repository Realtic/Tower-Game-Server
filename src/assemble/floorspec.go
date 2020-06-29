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

	return floorSpec
}

// AllFloors returns the floorspec for all floors
func (f *FloorSpecAssembler) AllFloors() error {
	// Read in the server side-saved account
	floors, err := f.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	f.Floors = &floors

	log.Print("successfully assembled floors")
	return nil
}
