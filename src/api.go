package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"tower/load"
	"tower/resapi"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/data/profile/{auth}", handleProfile).
		Queries("s", "{s}").
		Methods(http.MethodGet, http.MethodOptions)
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
	// TODO: move into resapi, meaning need to create func init for resapi rather than init inline
	writeStandardHeaders(w)

	writer := resapi.ResponseAPI{
		Writer: w,
	}

	var acc, err interface{}
	switch mux.Vars(r)["s"] {
	case "fresh":
		acc, err = load.Account(mux.Vars(r)["auth"], load.FRESH)
	case "stale":
		acc, err = load.Account(mux.Vars(r)["auth"], load.STALE)
	case "active":
		acc, err = load.Account(mux.Vars(r)["auth"], load.ACTIVE)
	default:
		writer.Error("invalid account status request")
		return
	}

	if err != nil {
		writer.Error("loading account failed")
		return
	}

	writer.PrintJSON(acc)
}
