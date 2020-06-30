package load

import (
	"log"

	"tower/assemble"
	datatype "tower/datatype/floorspec"
)

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
