package model

type ProjectActivity struct {
	ID              uint   `json:"id"`
	NumActivity     uint   `json:"numActivity"`
	ProjectActivity string `json:"projectActivity"`
	Category        string `json:"category"`
	NameProyect     string `json:"nameProyect"`
	NickUser        string `json:"nickUser"`
}

type ProjectActivities []ProjectActivity
