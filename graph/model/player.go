package model

type Player struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
	TeamID string  `json:"teamUd"`
	Team   *Team   `json:"team"`
}

type PlayerInput struct {
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
	TeamID string  `json:"teamId"`
}
