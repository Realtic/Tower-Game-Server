package synchronize

import (
	"log"
	"tower/calculate"
	"tower/datatype"
)

// SyncAccount synchronizes the given account up to the current time
// Mutates given account pointer
func SyncAccount(account *datatype.Account, timestamp int64) error {
	log.Print("begin of update account")

	if err := calculate.CalcTower(&(account.User.Tower), timestamp); err != nil {
		return err
	}

	log.Printf("saving lastsync to: %d", timestamp)
	account.User.Tower.LastSync = timestamp

	return nil
}
