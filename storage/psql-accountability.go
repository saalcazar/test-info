package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/ceadlbk-info/model"
)

type scannerAccountability interface {
	Scan(dest ...any) error
}

const (
	psqlCreateAccountability   = `SELECT create_accountability ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	psqlUpdateAccountability   = `SELECT update_accountability ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	psqlApprovedAccountability = `SELECT update_approved_accountability ($1, $2)`
	psqlDeleteAccountability   = `SELECT delete_accountability ($1)`
	psqlGetAccountabilityByID  = `SELECT id_accountability, presentation, amount, reception, cod_founder, name_proyect, signature, name_user, nick_user, id_activity, approved FROM accountabilities WHERE id_accountability = $1`
	psqlGetAllAccountability   = `SELECT id_accountability, presentation, amount, reception, cod_founder, name_proyect, signature, name_user, nick_user, id_activity, approved FROM accountabilities`
)

type PsqlAccountability struct {
	db *sql.DB
}

func NewPsqlAccountability(db *sql.DB) *PsqlAccountability {
	return &PsqlAccountability{db}
}

// Create
func (f *PsqlAccountability) Create(m *model.Accountability) error {
	stmt, err := f.db.Prepare(psqlCreateAccountability)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Amount,
		m.Reception,
		m.CodFounder,
		m.NameProyect,
		m.Signature,
		m.NameUser,
		m.NickUser,
		m.IdActivity,
		m.Approved,
	).Scan(&m.Reception)
	if err != nil {
		return err
	}
	fmt.Println("La rendición fue creada exitosamente")
	return nil
}

// Update
func (f *PsqlAccountability) Update(m *model.Accountability) error {
	stmt, err := f.db.Prepare(psqlUpdateAccountability)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Amount,
		&m.Reception,
		&m.CodFounder,
		&m.NameProyect,
		&m.Signature,
		&m.NameUser,
		&m.NickUser,
		&m.IdActivity,
		&m.Approved,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe la rendición con el id: %d", m.ID)
	}

	fmt.Println("La rendición se actualizó exitosamente")
	return nil
}

func (f *PsqlAccountability) Approved(m *model.Accountability) error {
	stmt, err := f.db.Prepare(psqlApprovedAccountability)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		&m.ID,
		&m.Approved,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe la rendición con el id: %d", m.ID)
	}

	fmt.Println("La rendición se aprobó exitosamente")
	return nil
}

// Delete
func (f *PsqlAccountability) Delete(id uint) error {
	stmt, err := f.db.Prepare(psqlDeleteAccountability)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("La rendición se eliminó correctamente")
	return nil
}

// GetByID
func (p *PsqlAccountability) GetByID(id uint) (*model.Accountability, error) {
	stmt, err := p.db.Prepare(psqlGetAccountabilityByID)
	if err != nil {
		return &model.Accountability{}, err
	}
	defer stmt.Close()
	return scanRowAccountabilities(stmt.QueryRow(id))
}

func (p *PsqlAccountability) GetAll() (model.Accountabilities, error) {
	stmt, err := p.db.Prepare(psqlGetAllAccountability)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(model.Accountabilities, 0)
	for rows.Next() {
		p, err := scanRowAccountabilities(rows)
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

func scanRowAccountabilities(s scannerAccountability) (*model.Accountability, error) {
	m := &model.Accountability{}
	err := s.Scan(
		&m.ID,
		&m.Presentation,
		&m.Amount,
		&m.Reception,
		&m.CodFounder,
		&m.NameProyect,
		&m.Signature,
		&m.NameUser,
		&m.NickUser,
		&m.IdActivity,
		&m.Approved,
	)
	if err != nil {
		return &model.Accountability{}, err
	}
	return m, nil
}
