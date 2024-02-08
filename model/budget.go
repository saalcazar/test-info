package model

type Budget struct {
	ID          uint   `json:"id"`
	Quantity    uint   `json:"quantity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	ImportUSD   uint   `json:"importUSD"`
	ImportBOB   uint   `json:"importBOB"`
	IdActivity  uint   `json:"idActivity"`
	CodFounder  string `json:"codFounder"`
	NickUser    string `json:"nickUser"`
}

type Budgets []Budget
