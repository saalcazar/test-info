package model

type Profile struct {
	ID      uint   `json:"id"`
	Profile string `json:"profile"`
	Nick    string `json:"nickUser"`
}

type Profiles []Profile
