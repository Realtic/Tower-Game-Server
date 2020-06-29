package calculate

import (
	"log"
	math "math/big"

	datatype_user "tower/datatype/user"
)

// TODO: Exp & Tower Level

// CalcTower recalculates the tower stats:
//  Cash from floors (Calculate what the current cash value should be, from last synchronize time & each floor income)
//  Exp from just Finished Construction/Upgrades (upgrades that finished since the last synchronize need to have exp added to current).
//  Level from Exp (Calculates what the level should be, given the current Exp), should be done after calculating new Exp.
func CalcTower(tower *datatype_user.Tower, currentTime int64) error {
	log.Print("Running calc for Tower")

	for i := 0; i < len(tower.Floors); i++ {
		if !isRentCollectable(&(tower.Floors[i])) {
			// Since it's not rent collectable, check if it can be upgraded (not opened...)
			// FloorOpened must be done manually by user by clicking on tower after it finishes upgrade
			upgradeIfUpgradable(&(tower.Floors[i]), tower.LastSync, currentTime)
			continue
		}

		log.Printf("can collect rent, thus old cash is: %d", tower.Cash)
		tower.Cash += calculateFloorIncome(&(tower.Floors[i]), tower.LastSync, currentTime)
		log.Printf("new cash is: %d", tower.Cash)
	}

	log.Print("Finished Calc")
	return nil
}

func isRentCollectable(floor *datatype_user.Floor) bool {
	return floor.FloorOpened
}

// Messy and probably not very efficient
func calculateFloorIncome(floor *datatype_user.Floor, lastTime int64, currentTime int64) int64 {
	hourly := new(math.Float).SetInt64(floor.HourlyRent)
	cTime := new(math.Float).SetInt64(currentTime)
	lTime := new(math.Float).SetInt64(lastTime)

	log.Printf("Last Time: %d, Current Time: %d, Hourly Rent: %d, Income Rate: %d",
		lTime, cTime, hourly, 3600)

	var total math.Float
	var tDiff math.Float
	var fHour math.Float

	// Equivalence: floor.HourlyRent * ((currentTime - lastTime) / 3600)
	fHour.SetInt64(3600)
	tDiff.Sub(cTime, lTime)
	tDiff.Quo(&tDiff, &fHour)
	total.Mul(hourly, &tDiff)

	totalInt, _ := total.Int64()

	return totalInt
}

// TODO:
func upgradeIfUpgradable(floor *datatype_user.Floor, lastTime int64, currentTime int64) {
	// if floor.UnderConstruction {
	// 	if
	// }
}
