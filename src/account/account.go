package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Account struct {
	user User `json:"user"`
}

type User struct {
	username string `json:"username"`
	password string `json:"password"`
	hash     string `json:"hash"`
	email    string `json:"email"`
	verified bool   `json:"verified"`
	deviceId string `json:"device-identifier"`
	tower    Tower  `json:"tower"`
}

type Tower struct {
	level  int     `json:"level"`
	exp    int     `json:"exp"`
	cash   int     `json:"cash"`
	gold   int     `json:"gold"`
	name   string  `json:"tower-name"`
	floors []Floor `json:"floors"`
}

type Floor struct {
	name           string `json:"name"`
	industry       string `json:"industry-type"`
	level          int    `json:"level"`
	currentWorkers int    `json:"current-workers"`
	maxWorkers     int    `json:"max-workers"`
	monthlyRent    int    `json:"monthly-rent"`
	construction   bool   `json:"under-construction"`
}

func Load(accountId string) (*Account, error) {
	var a Account

	jsonFile, err := os.Open("static/user.json")
	if err != nil {
		fmt.Println(err)
		return &a, err
	}

	defer jsonFile.Close()

	// TODO: Temporarily loads a static user file
	jsonBytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Print(err)
		return &a, err
	}

	if err := json.Unmarshal(jsonBytes, &a); err != nil {
		log.Print(err)
		return &a, err
	}

	fmt.Printf("The user is: %+v\n", a)
	// fmt.Printf("The tower level is: %d\n", a.user.tower.level)

	return &a, nil
}
