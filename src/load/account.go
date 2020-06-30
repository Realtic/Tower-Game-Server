package load

import (
	"errors"
	"log"

	cache "github.com/patrickmn/go-cache"

	"tower/assemble"
	datatype "tower/datatype/user"
)

// Type ...
type Type int

const (
	// STALE - Unsynchronized completely - last place left off
	STALE Type = iota
	// FRESH - Synchronized - but not written back to server
	FRESH
	// ACTIVE - Synchronized and written back to server
	ACTIVE
)

// TODO: Can merge these functions together

// Account assembles an account from the server
// Depending on the loadType - it'll load a stale, fresh or active account.
func Account(accountID string, loadType Type, floorCache *cache.Cache) (*datatype.Account, error) {
	assembler := assemble.InitAccount(accountID)
	if assembler.Error != nil {
		log.Print("error when calling assemble.InitAccount")
		return assembler.Account, assembler.Error
	}

	var err error
	switch loadType {
	case STALE:
		log.Print("assembling stale account")
		err = assembler.StaleAccount()
	case FRESH:
		log.Print("assembling fresh account")
		err = assembler.FreshAccount(floorCache)
	case ACTIVE:
		log.Print("assembling active account")
		err = assembler.ActiveAccount(floorCache)
	default:
		return nil, errors.New("invalid load type")
	}

	if err != nil {
		return assembler.Account, err
	}

	return assembler.Account, nil
}
