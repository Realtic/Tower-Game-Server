package datatype

// Floor data structure for a floor in a tower
type Floor struct {
	Name              string `json:"name"`
	Industry          string `json:"industry-type"`
	Level             int    `json:"level"`
	CurrentWorkers    int    `json:"current-workers"`
	MaxWorkers        int    `json:"max-workers"`
	MonthlyRent       int    `json:"monthly-rent"`
	Construction      bool   `json:"under-construction"`
	ConstructionStart int    `json:"construction-start"`
	FloorID           int    `json:"floor-id"`
}
