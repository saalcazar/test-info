package model

type User struct {
	ID           uint   `json:"id"`
	NameUser     string `json:"nameUser"`
	NickUser     string `json:"nickUser"`
	PasswordUser string `json:"passwordUser"`
	Charge       string `json:"charge"`
	Signature    string `json:"signature"`
	Profile      string `json:"profile"`
	Nick         string `json:"nick"`
	NameProyect  string `json:"nameProyect"`
}

type Users []User

type DataUser struct {
	Name        string
	Profile     string
	Signature   string
	NameProyect string
}
