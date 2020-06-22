package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tower/src/account"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot).Methods("GET")
	router.HandleFunc("/auth/login", handleLogin).Methods("POST")
	router.HandleFunc("/auth/register", handleRegister).Methods("POST")
	router.HandleFunc("/data/profile/{auth}", handleProfile).Methods("GET")

	log.Print("running")
	log.Fatal(http.ListenAndServe(":2468", router))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root api page")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]

	fmt.Fprintf(w, "Login\nemail: %s, password: %s", email, password)
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	email := mux.Vars(r)["email"]
	password := mux.Vars(r)["password"]

	fmt.Fprintf(w, "Register\nemail: %s, password: %s", email, password)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	acc, err := account.Load(mux.Vars(r)["auth"])
	if err != nil {
		fmt.Fprint(w, "error loading account")
		return
	}

	accountData, err := json.Marshal(acc)
	if err != nil {
		fmt.Fprint(w, "unable to marshal account json")
		return
	}

	fmt.Fprintf(w, "%s", string(accountData))
	// fmt.Fprintf(w, "%s", fileContents)
}
