package model

type Application struct {
	ID           uint   `json:"id"`
	Presentation string `json:"presentation"`
	Amount       uint   `json:"amount"`
	Approved     bool   `json:"approved"`
	NameProyect  string `json:"nameProyect"`
	Signature    string `json:"signature"`
	NameUser     string `json:"nameUSer"`
	NickUser     string `json:"nickUser"`
	IdActivity   uint   `json:"idActivity"`
}

type Applications []Application
