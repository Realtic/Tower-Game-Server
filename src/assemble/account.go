package assemble

import (
	"log"
	"time"
	"tower/datastore"
	"tower/datatype"
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
// Meaning that this funciton updates the server side datastore
// with the updated account details.
func (acc *AccountAssembler) FreshAccount() error {
	// Read in the server side-saved account
	account, err := acc.Store.Read()
	if err != nil {
		log.Print("error reading from account store")
		return err
	}

	acc.Account = &account

	// Synchronize the account
	err = synchronize.SyncAccount(acc.Account, acc.CurrentTime)
	if err != nil {
		log.Print("error when attempting to synchronize account")
		return err
	}

	// Write the new sync'd data back to server
	if err = acc.Store.Write(acc.Account); err != nil {
		log.Print("error writing to account store")
		return err
	}

	log.Print("successfully assembled fresh account")
	return nil
}

// TODO:
// // StaleAccount returns an assembled, unsynchronized account.
// // Meaning it does not update the server datastore. It's simply
// // a read-only transaction.
// func StaleAccount(accountID string) *AccountAssembler {
// 	// accountBlob := blob.AccountInit(accountID)

// 	// accountBlob.Read()
// 	// if accountBlob.Error != nil {
// 	// 	return nil, accountBlob.Error
// 	// }

// 	// return accountBlob.Account, nil

// 	return nil
// }