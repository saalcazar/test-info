package model

type Report struct {
	ID          uint   `json:"id"`
	Issues      string `json:"issues"`
	Results     string `json:"results"`
	Obstacle    string `json:"obstacle"`
	Conclusions string `json:"conclusions"`
	Anexos      string `json:"anexos"`
	Approved    bool   `json:"approved"`
	NameUser    string `json:"nameUser"`
	NameProyect string `json:"nameproyect"`
	Signature   string `json:"signature"`
	CodFounder  string `json:"codFounder"`
	NickUser    string `josn:"nickUser"`
	IdActivity  uint   `json:"idActivity"`
}

type Reports []Report
