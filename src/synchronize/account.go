package synchronize

import (
	"log"

	"tower/calculate"
	datatype "tower/datatype/user"
)

// Only resync if at least 3 seconds has passed sync tower last sync
// Mostly to save on system resources & b/c not really necessarily to resync often
var resyncMinTime = 3

// SyncAccount synchronizes the given account up to the current time
// Mutates given account pointer
func SyncAccount(account *datatype.Account, timestamp int64) error {
	log.Print("begin of account sync")

	if (timestamp - tower.LastSync) <= resyncMinTime {
		// Not an error, just warning
		log.Printf("premature resync, will not sync")
		return nil
	}

	if err := calculate.CalcTower(&(account.User.Tower), timestamp); err != nil {
		return err
	}

	log.Printf("saving lastsync to: %d", timestamp)
	account.User.Tower.LastSync = timestamp

	return nil
}
