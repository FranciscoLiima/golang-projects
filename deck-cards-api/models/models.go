package models

type CreateDeck struct {
	DeckID string `json:"deck_id"`
	Shuffled bool `json:"shuffled"`
	Remaining int `json:"remaining"`
}

type Shuffle struct {
	Shuffled bool `json:"shuffled"`
}

type Cards struct {
	//value string `json:"value"`
	//suit string `json:"suit"`
	Code string `json:"code"`
}

type OpenDeck struct {
	DeckID string `json:"deck_id"`
	Shuffled bool `json:"shuffled"`
	Remaining int `json:"remaining"`
	Cards []Cards `json:cards`
}
