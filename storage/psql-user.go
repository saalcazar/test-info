package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerUsers interface {
	Scan(dest ...any) error
}

const (
	psqlCreateUser  = `SELECT create_user ($1, $2, $3, $4, $5, $6, $7, $8)`
	psqlUpdateUser  = `SELECT update_user ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	psqlDeleteUser  = `SELECT delete_user ($1)`
	psqlGetUserByID = `SELECT id_user, name_user, nick_user, password_user, charge, signature, profile, nick, name_proyect FROM users WHERE id_user = $1`
	psqlGetAllUser  = `SELECT id_user, name_user, nick_user, password_user, charge, signature, profile, nick, name_proyect FROM users`
)

type PsqlUser struct {
	db *sql.DB
}

func NewPsqlUser(db *sql.DB) *PsqlUser {
	return &PsqlUser{db}
}

// Create
func (f *PsqlUser) Create(m *model.User) error {
	stmt, err := f.db.Prepare(psqlCreateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.NameUser,
		m.NickUser,
		m.PasswordUser,
		m.Charge,
		m.Signature,
		m.Profile,
		m.Nick,
		m.NameProyect,
	).Scan(&m.NameUser)
	if err != nil {
		return err
	}
	fmt.Println("El usuario fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlUser) Update(m *model.User) error {
	stmt, err := f.db.Prepare(psqlUpdateUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.NameUser,
		&m.NickUser,
		&m.PasswordUser,
		&m.Charge,
		&m.Signature,
		&m.Profile,
		&m.Nick,
		&m.NameProyect,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe un usuario con el id: %d", m.ID)
	}

	fmt.Println("El usuario se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlUser) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteUser)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El usuario se eliminó correctamente")
	return nil
}

// GetByID
func (p *PsqlUser) GetByID(id uint) (*model.User, error) {
	stmt, err := p.db.Prepare(psqlGetUserByID)
	if err != nil {
		return &model.User{}, err
	}
	defer stmt.Close()
	return scanRowUsers(stmt.QueryRow(id))
}

func (p *PsqlUser) GetAll() (model.Users, error) {
	stmt, err := p.db.Prepare(psqlGetAllUser)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Users, 0)
	for rows.Next() {
		p, err := scanRowUsers(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, *p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ms, nil
}

func scanRowUsers(s scannerUsers) (*model.User, error) {
	m := &model.User{}
	err := s.Scan(
		&m.ID,
		&m.NameUser,
		&m.NickUser,
		&m.PasswordUser,
		&m.Charge,
		&m.Signature,
		&m.Profile,
		&m.Nick,
		&m.NameProyect,
	)
	if err != nil {
		return &model.User{}, err
	}
	return m, nil
}
