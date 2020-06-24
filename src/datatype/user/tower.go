package user

// Tower data structure for a user's tower
type Tower struct {
	Level    int32   `json:"level"`
	Exp      int64   `json:"exp"`
	Cash     int64   `json:"cash"`
	Gold     int64   `json:"gold"`
	Name     string  `json:"tower-name"`
	LastSync int64   `json:"last-synchronized"`
	Floors   []Floor `json:"floors"`
}
