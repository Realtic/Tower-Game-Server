package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	cache "github.com/patrickmn/go-cache"

	"tower/load"
	"tower/resapi"
)

// AppCache store for the application
var floorCache *cache.Cache

func main() {
	if err := initFloorCache(); err != nil {
		log.Fatalf("unable to init floor cache; error: %s", err.Error())
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handleRoot).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/data/profile/{auth}", handleProfile).
		Queries("s", "{s}").
		Methods(http.MethodGet, http.MethodOptions)
	router.Use(mux.CORSMethodMiddleware(router))

	log.Print("running")
	log.Fatal(http.ListenAndServe(":2468", router))
}

func initFloorCache() error {
	expireTime := 5 * time.Minute
	purgeTime := 10 * time.Minute

	floorCache = cache.New(expireTime, purgeTime)
	// cache.DefaultExpiration, cache.NoExpiration

	const (
		prefix = "FLOOR-CACHE_"
	)

	// Loads all floors and saves each one individually to cache
	log.Print("saving all floors")

	allFloors, err := load.AllFloors()
	if err != nil {
		log.Print("unable to load all floors")
		return err
	}

	for i := 0; i < len(allFloors.Floors); i++ {
		log.Print("floor need to save to cache")

		json, err := json.Marshal(allFloors.Floors[i])
		if err != nil {
			log.Fatal("unable to marshal floor for writing son to floorCache, error: %s", err.Error())
		}

		floorStringID := prefix + strconv.Itoa(allFloors.Floors[i].FloorID)
		log.Printf("the floor string id when saving floorcache: %s", floorStringID)

		floorCache.Set(floorStringID, json, cache.NoExpiration)
	}

	return nil
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
		acc, err = load.Account(mux.Vars(r)["auth"], load.FRESH, floorCache)
	case "stale":
		acc, err = load.Account(mux.Vars(r)["auth"], load.STALE, floorCache)
	case "active":
		acc, err = load.Account(mux.Vars(r)["auth"], load.ACTIVE, floorCache)
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
