package user

// Floor data structure for a floor in a tower
type Floor struct {
	Name              string `json:"name"`
	FloorID           int    `json:"floor-id"`
	Industry          string `json:"industry-type"`
	Level             int    `json:"level"`
	CurrentWorkers    int    `json:"current-workers"`
	MaxWorkers        int    `json:"max-workers"`
	MonthlyRent       int64  `json:"monthly-rent"`
	UnderConstruction bool   `json:"under-construction"`
	ConstructionStart int    `json:"construction-start"`
	FloorOpened       bool   `json:"floor-opened"`
}
