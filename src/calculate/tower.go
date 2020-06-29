package calculate

import (
	"log"

	datatype "tower/datatype/user"
)

// TODO: Exp & Tower Level
// CalcTower recalculates the tower stats:
//  Cash from floors (Calculate what the current cash value should be, from last synchronize time & each floor income)
//  Exp from just Finished Construction/Upgrades (upgrades that finished since the last synchronize need to have exp added to current).
//  Level from Exp (Calculates what the level should be, given the current Exp), should be done after calculating new Exp.
func CalcTower(tower *datatype.Tower, currentTime int64) error {
	log.Print("Running calc for Tower")

	for i := 0; i < len(tower.Floors); i++ {
		if !isRentCollectable(&(tower.Floors[i])) {
			// Since it's not rent collectable, check if it can be upgraded (not opened...)
			// FloorOpened must be done manually by user by clicking on tower after it finishes upgrade
			upgradeIfUpgradable(&(tower.Floors[i]), tower.LastSync, currentTime)
			continue
		}

		tower.Cash += calculateFloorIncome(&(tower.Floors[i]), tower.LastSync, currentTime)
	}

	log.Print("Finished Calc")
	return nil
}

func isRentCollectable(floor *datatype.Floor) bool {
	return floor.FloorOpened
}

func calculateFloorIncome(floor *datatype.Floor, lastTime int64, currentTime int64) int64 {
	return floor.MonthlyRent * ((currentTime - lastTime) / IncomeRate)
}

// TODO:
func upgradeIfUpgradable(floor *datatype.Floor, lastTime int64, currentTime int64) {
	// if floor.UnderConstruction {
	// 	if
	// }
}
