package load

import (
	"log"

	"tower/assemble"
	datatype "tower/datatype/floorspec"
)

// TODO: load in all floors to app memory when the app starts,
// Thus, instead of reading from file every time - read only once into memory when app starts...
// Can do this since floorspec.json isn't expected to really change while the app is running.

// AllFloors loads all the floors
func AllFloors() (*datatype.Floors, error) {
	assembler := assemble.InitFloorSpec()
	if assembler.Error != nil {
		log.Print("error when calling assemble.InitFloorSpec")
		return assembler.Floors, assembler.Error
	}

	err := assembler.AllFloors()
	if err != nil {
		return assembler.Floors, err
	}

	log.Print("successfully loaded all floors")
	return assembler.Floors, nil
}
