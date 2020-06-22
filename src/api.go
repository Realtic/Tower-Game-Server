package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tower/load"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/data/profile/{auth}", handleProfile).Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))

	log.Print("running")
	log.Fatal(http.ListenAndServe(":2468", router))
}

func writeStandardHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	writeStandardHeaders(w)

	response := map[string]string{
		"response": "root api page",
	}

	json, err := json.Marshal(response)
	if err != nil {
		log.Print(err)
		return
	}

	fmt.Fprint(w, string(json))
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	writeStandardHeaders(w)

	acc, err := load.AccountFresh(mux.Vars(r)["auth"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error loading account")
		return
	}

	accountData, err := json.Marshal(acc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "unable to marshal account json")
		return
	}

	fmt.Fprintf(w, "%s", string(accountData))
}
