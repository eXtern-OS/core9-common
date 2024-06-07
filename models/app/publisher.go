package app

type Publisher struct {
	Name        string  `json:"name"`
	ActiveSince int64   `json:"active_since"` // Timestamp of first activity
	Rating      float64 `json:"rating"`       //rating across all apps
	IsVerified  bool    `json:"is_verified"`
}
