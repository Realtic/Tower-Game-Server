package datatype

// Tower data structure for a user's tower
type Tower struct {
	Level    int     `json:"level"`
	Exp      int     `json:"exp"`
	Cash     int     `json:"cash"`
	Gold     int     `json:"gold"`
	Name     string  `json:"tower-name"`
	LastSync int64   `json:"last-synchronized"`
	Floors   []Floor `json:"floors"`
}
