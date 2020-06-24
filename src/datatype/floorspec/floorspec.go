package floorspec

// Floors ...
type Floors struct {
	Floors []Floor `json:"floors"`
}

// Floor ...
type Floor struct {
	DefaultName               string       `json:"default-name"`
	Industry                  string       `json:"industry-type"`
	FloorID                   int          `json:"floor-id"`
	InitialConstructionSprite string       `json:"initial-construction-sprite"`
	ConstructionCost          int          `json:"construction-cost"`
	ConstructionTime          int          `json:"construction-time"`
	FloorLevels               []FloorLevel `json:"floor-levels"`
}

// FloorLevel ...
type FloorLevel struct {
	Level         int    `json:"level"`
	UpgradeTime   int    `json:"upgrade-time"`
	UpgradeCost   int    `json:"upgrade-cost"`
	Income        int    `json:"income"`
	MaxWorkers    int    `json:"max-workers"`
	Sprite        string `json:"sprite"`
	UpgradeSprite string `json:"upgrade-sprite"`
}
