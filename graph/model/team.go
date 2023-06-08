package model

type Team struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Country string    `json:"country"`
	Players []*Player `json:"players"`
}

type TeamInput struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}
