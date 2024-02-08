package model

type Founder struct {
	ID          uint   `json:"id"`
	CodFounder  string `json:"codFounder"`
	NameFounder string `json:"nameFounder"`
	NickUser    string `json:"nickUser"`
}

type Founders []Founder
