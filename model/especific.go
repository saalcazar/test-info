package model

type Especific struct {
	ID           uint   `json:"id"`
	NumEspecific uint   `json:"numEspecific"`
	Especific    string `json:"especific"`
	NickUser     string `json:"nickUser"`
	NameProyect  string `json:"nameProyect"`
}

type Especifics []Especific
