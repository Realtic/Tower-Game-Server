package calculate

import (
	"log"
	"tower/datatype"
)

// CalcTower recalculates the tower stats:
//  Cash from floors (Calculate what the current cash value should be, from last synchronize time & each floor income)
//  Exp from just Finished Construction/Upgrades (upgrades that finished since the last synchronize need to have exp added to current).
//  Level from Exp (Calculates what the level should be, given the current Exp), should be done after calculating new Exp.
func CalcTower(tower *datatype.Tower, timestamp int64) error {
	log.Print("Running calc for Tower")

	for i := 0; i < len(tower.Floors); i++ {
		log.Print("Running calc for a Floor")
		CalcFloor(&(tower.Floors[i]), timestamp)
	}

	log.Print("Finished Calc")
	return nil
}
