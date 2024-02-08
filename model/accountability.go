package model

type Accountability struct {
	ID           uint   `json:"id"`
	Presentation string `json:"presentation"`
	Amount       uint   `json:"amount"`
	Reception    string `json:"reception"`
	CodFounder   string `json:"codFounder"`
	NameProyect  string `json:"nameProyect"`
	Signature    string `json:"signature"`
	NameUser     string `json:"nameUser"`
	NickUser     string `json:"nickUser"`
	IdActivity   uint   `json:"idActivity"`
	Approved     bool   `json:"approved"`
}

type Accountabilities []Accountability
