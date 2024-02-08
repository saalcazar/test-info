package model

type Quantitative struct {
	ID         uint   `json:"id"`
	Achieved   uint   `json:"achieved"`
	Day        string `json:"day"`
	SpFemale   uint   `json:"spFemale"`
	SpMale     uint   `json:"spMale"`
	FFemale    uint   `json:"fFemale"`
	FMale      uint   `json:"fMale"`
	NaFemale   uint   `json:"naFemale"`
	NaMale     uint   `json:"naMale"`
	PFemale    uint   `json:"pFemale"`
	PMale      uint   `json:"pMale"`
	IdActivity uint   `json:"idActivity"`
	NickUser   string `json:"nickUser"`
}

type Quantitatives []Quantitative
