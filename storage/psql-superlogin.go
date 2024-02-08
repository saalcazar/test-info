package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

const (
	psqlSuperLoginUser = `SELECT validate_super_user ($1, $2)`
)

type PsqlSuperLogin struct {
	db *sql.DB
}

func NewPsqlSuperLogin(db *sql.DB) *PsqlSuperLogin {
	return &PsqlSuperLogin{db}
}

// Create
func (pr *PsqlSuperLogin) SuperLogin(m *model.SuperLogin) bool {
	stmt, err := pr.db.Prepare(psqlSuperLoginUser)
	if err != nil {
		return false
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.NickUser,
		m.Password,
	).Scan(&m.NickUser)
	if err != nil {
		return false
	}
	fmt.Println("El perfil fue validado exitosamente")
	return true
}
