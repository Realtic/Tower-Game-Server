package assemble

import (
	"log"
	"tower/datastore"
	datatype "tower/datatype/floorspec"
)

// FloorSpecAssembler ...
type FloorSpecAssembler struct {
	Store  *datastore.FloorSpecStore
	Floors *datatype.Floors
	Error  error
}

// InitFloorSpec ...
func InitFloorSpec() *FloorSpecAssembler {
	var floorSpec = new(FloorSpecAssembler)
	floorSpec.Store = datastore.InitFloorSpecStore()

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
