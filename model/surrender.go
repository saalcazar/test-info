package model

type Surrender struct {
	ID            uint   `json:"id"`
	DateInvoice   string `json:"dateInvoice"`
	InvoiceNumber string `json:"invoiceNumber"`
	Code          string `json:"code"`
	Description   string `json:"description"`
	ImportUSD     uint   `json:"importUSD"`
	ImportBOB     uint   `json:"importBOB"`
	IdActivity    uint   `json:"idActivity"`
	NickUser      string `json:"nickUser"`
}

type Surrenders []Surrender
