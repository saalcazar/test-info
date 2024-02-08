package model

type ProjectResult struct {
	ID            uint   `json:"id"`
	NumResult     uint   `json:"numResult"`
	ProjectResult string `json:"projectResult"`
	NameProyect   string `json:"nameProyect"`
	NickUser      string `json:"nickUser"`
}

type ProjectResults []ProjectResult
