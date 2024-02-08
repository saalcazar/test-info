package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerFounder interface {
	Scan(dest ...any) error
}

const (
	psqlCreateFounder  = `SELECT create_founder ($1, $2, $3)`
	psqlUpdateFounder  = `SELECT update_founder ($1, $2, $3, $4)`
	psqlDeleteFounder  = `SELECT delete_founder ($1)`
	psqlGetFounderByID = `SELECT id_founder, cod_founder, name_founder, nick_user FROM founders WHERE id_founder = $1`
	psqlGetAllFounder  = `SELECT id_founder, cod_founder, name_founder, nick_user FROM founders`
)

type PsqlFounder struct {
	db *sql.DB
}

func NewPsqlFounder(db *sql.DB) *PsqlFounder {
	return &PsqlFounder{db}
}

// Create
func (f *PsqlFounder) Create(m *model.Founder) error {
	stmt, err := f.db.Prepare(psqlCreateFounder)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.CodFounder,
		m.NameFounder,
		m.NickUser,
	).Scan(&m.CodFounder)
	if err != nil {
		return err
	}
	fmt.Println("El financiador fue creado exitosamente")
	return nil
}

// Update
func (f *PsqlFounder) Update(m *model.Founder) error {
	stmt, err := f.db.Prepare(psqlUpdateFounder)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.CodFounder,
		&m.NameFounder,
		&m.NickUser,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe el financiador con el id: %d", m.ID)
	}

	fmt.Println("El financiador se actualizó exitosamente")
	return nil
}

// Delete
func (f *PsqlFounder) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteFounder)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("El financiador se eliminó correctamente")
	return nil
}

func (p *PsqlFounder) GetByID(id uint) (*model.Founder, error) {
	stmt, err := p.db.Prepare(psqlGetFounderByID)
	if err != nil {
		return &model.Founder{}, err
	}
	defer stmt.Close()
	return scanRowFounders(stmt.QueryRow(id))
}

func (p *PsqlFounder) GetAll() (model.Founders, error) {
	stmt, err := p.db.Prepare(psqlGetAllFounder)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Founders, 0)
	for rows.Next() {
		p, err := scanRowFounders(rows)
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

func scanRowFounders(s scannerFounder) (*model.Founder, error) {
	m := &model.Founder{}
	err := s.Scan(
		&m.ID,
		&m.CodFounder,
		&m.NameFounder,
		&m.NickUser,
	)
	if err != nil {
		return &model.Founder{}, err
	}
	return m, nil
}
