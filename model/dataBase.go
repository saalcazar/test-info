package model

type DataBase struct {
	ID               uint   `json:"id"`
	NameParticipant  string `json:"nameParticipant"`
	Gender           string `json:"gender"`
	Age              uint   `json:"age"`
	Organization     string `json:"organization"`
	Phone            string `json:"phone"`
	TypeParticipant  string `json:"typeParticipant"`
	NameProyect      string `json:"nameProyect"`
	CodFounder       string `json:"codFounder"`
	Activity         string `json:"activity"`
	NickUser         string `json:"nickUser"`
	Municipality     string `json:"municipality"`
	TypeOrganization string `json:"typeOrganization"`
}

type DataBases []DataBase
