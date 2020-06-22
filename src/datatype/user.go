package datatype

// User data structure for an account's user
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	DeviceID string `json:"device-identifier"`
	Tower    Tower  `json:"tower"`
}
