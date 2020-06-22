package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tower/account"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot).Methods("GET")
	router.HandleFunc("/data/profile/{auth}", handleProfile).Methods("GET")

	log.Print("running")
	account.Load("test")
	log.Fatal(http.ListenAndServe(":2468", router))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "root api page")
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
}
