package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerProfile interface {
	Scan(dest ...any) error
}

const (
	psqlCreateProfile = `SELECT create_profile ($1, $2)`
	psqlUpdateProfile = `SELECT update_profile ($1, $2, $3)`
	psqlDeleteProfile = `SELECT delete_profile ($1)`
	psqlGetAllProfile = `SELECT id_profile, profile, nick FROM profiles`
)

type PsqlProfile struct {
	db *sql.DB
}

func NewPsqlProfile(db *sql.DB) *PsqlProfile {
	return &PsqlProfile{db}
}

// Create
func (pr *PsqlProfile) Create(m *model.Profile) error {
	stmt, err := pr.db.Prepare(psqlCreateProfile)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Profile,
		m.Nick,
	).Scan(&m.Profile)
	if err != nil {
		return err
	}
	fmt.Println("El perfil fue creado exitosamente")
	return nil
}

// Update
func (pr *PsqlProfile) Update(m *model.Profile) error {
	stmt, err := pr.db.Prepare(psqlUpdateProfile)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Profile,
		&m.Nick,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe un perfil con el id: %d", m.ID)
	}

	fmt.Println("El perfil se actualizó exitosamente")
	return nil
}

// Delete
func (pr *PsqlProfile) Delete(id uint) error {
	stmt, err := pr.db.Prepare(psqlDeleteProfile)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El perfil se eliminó correctamente")
	return nil
}

func (p *PsqlProfile) GetAll() (model.Profiles, error) {
	stmt, err := p.db.Prepare(psqlGetAllProfile)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Profiles, 0)
	for rows.Next() {
		p, err := scanRowProfiles(rows)
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

func scanRowProfiles(s scannerProfile) (*model.Profile, error) {
	m := &model.Profile{}
	err := s.Scan(
		&m.ID,
		&m.Profile,
		&m.Nick,
	)
	if err != nil {
		return &model.Profile{}, err
	}
	return m, nil
}
