package model

type Proyect struct {
	ID          uint   `json:"id"`
	CodProyect  string `json:"codProyect"`
	NameProyect string `json:"nameProyect"`
	Objetive    string `json:"objetive"`
	CodFounder  string `json:"codFounder"`
	NickUser    string `json:"nickUser"`
}

type Proyects []Proyect
