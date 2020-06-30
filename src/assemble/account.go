package assemble

import (
	"log"
	"time"

	cache "github.com/patrickmn/go-cache"

	"tower/datastore"
	datatype "tower/datatype/user"
	"tower/synchronize"
)

// AccountAssembler ...
type AccountAssembler struct {
	AccountID   string
	Fresh       bool // fresh or stale
	Store       *datastore.AccountStore
	Account     *datatype.Account
	Error       error
	CurrentTime int64
}

// InitAccount ...
func InitAccount(accountID string) *AccountAssembler {
	var account = new(AccountAssembler)
	account.Store = datastore.InitAccountStore(accountID)
	account.AccountID = accountID
	account.CurrentTime = time.Now().Unix()

	return account
}

// FreshAccount returns an assembled, synchronized account.
func (acc *AccountAssembler) FreshAccount(floorCache *cache.Cache) error {
	// Read in the server side-saved account
	account, err := acc.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	acc.Account = &account

	// Synchronize the account
	err = synchronize.SyncAccount(acc.Account, acc.CurrentTime, floorCache)
	if err != nil {
		log.Print("error when attempting to synchronize account")
		return err
	}

	log.Print("successfully assembled fresh account")
	return nil
}

// ActiveAccount returns a FreshAccount that's also written back to the server
func (acc *AccountAssembler) ActiveAccount(floorCache *cache.Cache) error {
	// Read in the server side-saved account
	account, err := acc.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	acc.Account = &account

	// Synchronize the account
	err = synchronize.SyncAccount(acc.Account, acc.CurrentTime, floorCache)
	if err != nil {
		log.Print("error when attempting to synchronize account")
		return err
	}

	// Write the new sync'd data back to server
	if err = acc.Store.Write(acc.Account); err != nil {
		log.Print("error writing to account store")
		return err
	}

	log.Print("successfully assembled active account")
	return nil
}

// StaleAccount returns an assembled, unsynchronized account.
// Meaning it does not update the server datastore. It's simply
// a read-only transaction.
func (acc *AccountAssembler) StaleAccount() error {
	// Read in the server side-saved account
	account, err := acc.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	acc.Account = &account

	log.Print("successfully assembled stale account")
	return nil
}
