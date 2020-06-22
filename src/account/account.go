package account

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Floor data structure for a floor in a tower
type Floor struct {
	Name           string `json:"name"`
	Industry       string `json:"industry-type"`
	Level          int    `json:"level"`
	CurrentWorkers int    `json:"current-workers"`
	MaxWorkers     int    `json:"max-workers"`
	MonthlyRent    int    `json:"monthly-rent"`
	Construction   bool   `json:"under-construction"`
}

// Tower data structure for a user's tower
type Tower struct {
	Level  int     `json:"level"`
	Exp    int     `json:"exp"`
	Cash   int     `json:"cash"`
	Gold   int     `json:"gold"`
	Name   string  `json:"tower-name"`
	Floors []Floor `json:"floors"`
}

// User data structure for an account's user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	DeviceID string `json:"device-identifier"`
	Tower    Tower  `json:"tower"`
}

// Account data structure
type Account struct {
	User User `json:"user"`
}

// Load the specific accountId's account into memory
func Load(accountID string) (*Account, error) {
	var a Account

	// TODO: Temporarily loads a static user file
	jsonBytes, err := ioutil.ReadFile("../static/user.json")
	if err != nil {
		log.Print(err)
		return &a, err
	}

	if err := json.Unmarshal(jsonBytes, &a); err != nil {
		log.Print(err)
		return &a, err
	}

	log.Printf("The user is: %+v\n", a)

	return &a, nil
}
