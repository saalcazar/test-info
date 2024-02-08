package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

const (
	psqlLoginUser = `SELECT name_user, profile, signature, name_proyect FROM validate_user ($1, $2)`
)

type PsqlLogin struct {
	db *sql.DB
}

func NewPsqlLogin(db *sql.DB) *PsqlLogin {
	return &PsqlLogin{db}
}

// Create
func (pr *PsqlLogin) Login(m *model.Login) (bool, []*model.DataUser, error) {
	stmt, err := pr.db.Prepare(psqlLoginUser)
	if err != nil {
		return false, nil, err
	}
	defer stmt.Close()

	var nameUser, profile, signature, nameProyect string

	err = stmt.QueryRow(
		m.NickUser,
		m.Password,
	).Scan(&nameUser, &profile, &signature, &nameProyect)
	if err != nil {
		fmt.Printf("Error fatal:  %v\n", err)
		return false, nil, err
	}

	user := &model.DataUser{
		Name:        nameUser,
		Profile:     profile,
		Signature:   signature,
		NameProyect: nameProyect,
	}

	fmt.Println("El perfil fue validado exitosamente")
	fmt.Println(stmt.QueryRow())
	return true, []*model.DataUser{user}, err
}
