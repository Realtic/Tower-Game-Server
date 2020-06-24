package load

import (
	"log"

	"tower/assemble"
	datatype "tower/datatype/user"
)

// TODO: Can merge these functions together

// FreshAccount returns an assembled (synchronized-server) account
// Writes the updated account back to server
func FreshAccount(accountID string) (*datatype.Account, error) {
	assembler := assemble.InitAccount(accountID)
	if assembler.Error != nil {
		log.Print("error when calling assemble.InitAccount")
		return assembler.Account, assembler.Error
	}

	err := assembler.FreshAccount()
	if err != nil {
		return assembler.Account, err
	}

	log.Print("successfully loaded fresh account")
	return assembler.Account, nil
}

// StaleAccount returns an assembled (unsynchronized-server) account
// Is read only, doesn't update the server with the account's newest details
func StaleAccount(accountID string) (*datatype.Account, error) {
	assembler := assemble.InitAccount(accountID)
	if assembler.Error != nil {
		log.Print("error when calling assemble.InitAccount")
		return assembler.Account, assembler.Error
	}

	err := assembler.StaleAccount()
	if err != nil {
		return assembler.Account, err
	}

	log.Print("successfully loaded stale account")
	return assembler.Account, nil
}
