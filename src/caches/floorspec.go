package caches

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"

	cache "github.com/patrickmn/go-cache"

	datatype_floorspec "tower/datatype/floorspec"
)

const (
	prefix = "FLOOR-CACHE_"
)

// GetFloor gets the specified floorID floor from the floorcache.
func GetFloor(floorID int, floorCache *cache.Cache) (*datatype_floorspec.Floor, error) {
	var floor datatype_floorspec.Floor

	log.Print("saving all floors")

	floorInterface, found := floorCache.Get(prefix + strconv.Itoa(floorID))
	if !found {
		errorStr := fmt.Sprintf("unable to get floor from cache: ('%s')", prefix+strconv.Itoa(floorID))
		return &floor, errors.New(errorStr)
	}

	floorBytes, worked := floorInterface.([]byte)
	if !worked {
		log.Print("type assertion to []byte failed")
		return &floor, errors.New("type assertion to []byte failed")
	}

	log.Printf("success, the floor is: %s", floorBytes)

	if err := json.Unmarshal(floorBytes, &floor); err != nil {
		log.Print(err)
		return &floor, err
	}

	return &floor, nil
}
