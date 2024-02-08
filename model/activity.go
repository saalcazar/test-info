package model

type Activity struct {
	ID              uint   `json:"id"`
	Activity        string `json:"activity"`
	DateStar        string `json:"dateStar"`
	DateEnd         string `json:"dateEnd"`
	Place           string `json:"place"`
	Expected        uint   `json:"expected"`
	Objetive        string `json:"objetive"`
	ResultExpected  string `json:"resultExpected"`
	Description     string `json:"description"`
	NameProyect     string `json:"nameProyect"`
	CodFounder      string `json:"codFounder"`
	Especific       string `json:"especific"`
	NickUser        string `json:"nickUser"`
	ProjectResult   string `json:"projectResult"`
	ProjectActivity string `json:"projectActivity"`
}

type Activities []Activity
