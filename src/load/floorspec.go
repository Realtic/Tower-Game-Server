package load

import (
	"log"

	"tower/assemble"
	datatype "tower/datatype/floorspec"
)

// TODO: specific floor load

// AllFloors ...
func AllFloors(floorID int) (*datatype.Floors, error) {
	assembler := assemble.InitFloorSpec(floorID)
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
